from __future__ import unicode_literals

import json
import spider
import sql
from multiprocessing import Process

tags = [
    "喜剧",
    "生活",
    "爱情",
    "动作",
    "科幻",
    "悬疑",
    "惊悚",
    "动画",
    "奇幻",
]


def spider_subject(t, db, start, batch_size):
    cnt = 0
    while True:
        try:
            data, size = spider.get_subjects(t, batch_size, start)
            cnt += size
            if size <= 0:
                break
            start += batch_size
        except BaseException:
            print("failed to get subjects list")
            break

        for subject in data:
            mid = subject["mid"]
            name = subject["title"]
            avatar = subject["avatar"]

            # 懒得分析原因了，一把唆
            try:

                plot, rates, detail = spider.get_subject_info(mid)

                score = json.dumps(
                    {"one": rates["1星"], "two": rates["2星"], "three": rates["3星"], "four": rates["4星"],
                     "five": rates["5星"],
                     "score": subject["rate"], "total_cnt": int(rates["总数"])})

                if detail["上映日期"].__contains__("("):
                    date = detail["上映日期"][:detail["上映日期"].index("(")] + " 00:00:00"
                else:
                    date = detail["上映日期"] + " 00:00:00"

                tags = detail["标签"]
                stars = int(float(subject["rate"]) / 2)
                details = dict()
                celebrities = []
                for cid in {**detail["导演"], **detail["编剧"], **detail["主演"]}.keys():
                    celebrities.append(int(cid))
                celebrities = json.dumps(celebrities)

                details["director"] = ",".join(detail["导演"].values())
                details["type"] = tags.split(",")
                details["IMDb"] = detail["IMDb"]
                details["period"] = int(detail["片长"])
                details["region"] = detail["制片国家/地区"]
                details["release"] = date
                details["website"] = detail.setdefault("官方网站", "")
                details["writers"] = list(detail["编剧"].values())
                details["language"] = detail["语言"]
                details["nicknames"] = detail["又名"].split(" / ")
                details["characters"] = list(detail["主演"].values())
                details = json.dumps(details, ensure_ascii=False)

                print(f"prepare add {name}")
                db.insert_subject(mid=mid, tags=tags, date=date, stars=stars, detail=details, name=name, score=score,
                                  plot=plot, avatar=avatar, celebrities=celebrities)
                print(f"added a new movie! mid: {mid} name: {name}")
            except BaseException:
                print(f"failed adding the movie: {name}")

    print(f"tag: {t} total data size: {cnt}, prepare to close db")
    db.close()


# 喜剧 260
# 多进程爬取
# 不过要注意豆瓣可能会ban异常ip
def subject_insert(tags_param=None):
    if tags_param is None:
        tags_param = [
            "喜剧",
            "生活",
            "爱情",
            "动作",
            "科幻",
            "悬疑",
            "惊悚",
            "动画",
            "奇幻",
        ]
    lp = []
    for tag in tags_param:
        print(f"add movie of {tag}")
        p = Process(target=spider_subject, args=(tag, sql.DB(), 0, 20,))
        p.start()
        lp.append(p)

    map(lambda i: i.join(), lp)


def spider_celebrity(db, start, batch_size):
    while True:
        celebrities_tuple = db.fetch_all(f"SELECT celebrities FROM subject LIMIT {batch_size} OFFSET {start}")
        if len(celebrities_tuple) <= 0:
            break

        start += batch_size

        for celebrities in celebrities_tuple:
            celebrities = json.loads(celebrities[0])
            for cid in celebrities:
                try:
                    name, name_en, avatar, detail, brief = spider.get_celebrity_info(cid)
                    print(f"prepare add {name}")
                    db.insert_celebrity(
                        cid=cid,
                        name=name,
                        name_en=name_en,
                        avatar=avatar,
                        gender=detail.setdefault("性别", "未知"),
                        sign=detail.setdefault("星座", "未知"),
                        birth=detail.setdefault("出生日期", "未知"),
                        hometown=detail.setdefault("出生地", "未知"),
                        job=detail.setdefault("职业", "未知"),
                        imdb=detail.setdefault("imdb编号", "未知"),
                        brief=brief
                    )
                    print(f"added celebrity: {cid}, {name}")
                except BaseException:
                    print(f"failed add celebrity : {cid}")
                    continue

    print("prepare close db")
    db.close()


# 豆瓣会封ip...
if __name__ == '__main__':
    # pl = []
    # for i in range(7,8):
    #     print(f"start at {155*i}, end at {155*i+155}")
    #     p = Process(target=spider_celebrity, args=(sql.DB(), 1200, 40,))
    #     p.start()
    #     pl.append(p)
    # map(lambda c: c.join(), pl)
    pass
    #subject_insert(["生活"])
