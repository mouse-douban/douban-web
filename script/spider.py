import re
import requests
from bs4 import BeautifulSoup

# 快乐爬虫

subjects_api = "https://movie.douban.com/j/search_subjects?type=movie&tag={tag}&sort=rank&page_limit={" \
               "page_limit}&page_start={page_start}"

subject_info_url = "https://movie.douban.com/subject/{mid}"

celebrity_info_url = "https://movie.douban.com/celebrity/{cid}"


def get_subjects(tag, limit, start):
    limit = str(limit)
    start = str(start)
    url = subjects_api.replace("{tag}", tag).replace("{page_limit}", limit).replace("{page_start}", start)
    print(f"request {url}")
    resp = requests.get(url=url, headers={"Referer": "movie.douban.com",
                                          "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) "
                                                        "AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.2 "
                                                        "Safari/605.1.15"})
    resp_json = resp.json()["subjects"]

    return list(map(lambda i: {"rate": i["rate"], "title": i["title"], "mid": i["id"], "avatar": i["cover"]}, resp_json)), len(resp_json)


'''
test

rate, cnt = get_subjects("恐怖", "20", "0")
print(rate, cnt)
'''


def get_subject_info(mid):
    mid = str(mid)
    url = subject_info_url.replace("{mid}", mid)
    resp = requests.get(url=url, headers={"Referer": "movie.douban.com",
                                          "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.2 Safari/605.1.15"})
    bs = BeautifulSoup(resp.text, "lxml").find("div", id="content")
    w_rates = bs.find("div", class_="ratings-on-weight").find_all("div", class_="item")

    plot = bs.find('span', class_="", property="v:summary").text
    rates = {"总数": bs.find("span", property="v:votes").text}
    detail = {}

    for rate in w_rates:
        rate_star = str(rate.find("span").text).strip()
        rate_per = rate.find("span", class_="rating_per").text
        rates[rate_star] = rate_per

    subject_detail = bs.find("div", class_="subject clearfix").find("div", id="info")
    detail["标签"] = str(",").join(map(lambda tag: str(tag.text), subject_detail.find_all("span", property="v:genre")))

    for pl in subject_detail.find_all("span", class_="pl"):
        if not ["制片国家/地区:", "语言:", "又名:", "IMDb:", "官方网站:"].__contains__(pl.text):
            continue

        start = str(subject_detail).index(pl.text)
        end = str(subject_detail).index("<br/>", start)
        substr = str(subject_detail)[start + len(pl.text) + len("</span>"):end].strip()
        detail[str(pl.text).strip(":")] = substr
        if str(pl.text) == "官方网站:":
            detail[str(pl.text).strip(":")] = subject_detail.find("a", rel="nofollow").text

    detail["上映日期"] = subject_detail.find("span", property="v:initialReleaseDate").text
    detail["片长"] = subject_detail.find("span", property="v:runtime").text.replace("分钟", "")
    if detail["片长"].__contains__("("):
        detail["片长"] = detail["片长"][:detail["片长"].index("(")]

    for k, attrs in zip(["导演", "编剧", "主演"], subject_detail.find_all("span", class_="attrs")[:3]):
        celebrities = {}
        for attr in filter(lambda a: str(a.get("href")).__contains__("celebrity"), attrs.find_all("a")):
            celebrities[str(attr.get("href")).lstrip("/celebrity/").rstrip("/")] = attr.text
        detail[k] = celebrities

    return re.sub(r"\s+", "", plot), rates, detail


'''
test

plot, rates, detail = get_subject_info("1292225")
print(plot, rates, detail)
'''


def get_celebrity_info(cid):
    cid = str(cid)
    url = celebrity_info_url.replace("{cid}", cid)
    resp = requests.get(url=url, headers={"Referer": "movie.douban.com",
                                          "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.2 Safari/605.1.15"})
    bs = BeautifulSoup(resp.text, "lxml").find("div", id="content")

    name, name_en = bs.find("h1").text.split(" ")[0], bs.find("h1").text.split(" ")[1]

    headline = bs.find("div", id="headline", class_="item")
    avatar = headline.find("img").get("src")
    detail = {}
    brief = bs.find("div", class_="bd").text.strip()
    for i in headline.find("div", class_="info").find_all("li"):
        key = i.find("span").text
        detail[key] = str(i.text).replace("<span>", "").replace("</span>", "").replace(key, "").strip().strip(" :").strip()

    return name, name_en, avatar, detail, brief


'''
test

name, name_en, avatar, detail, brief = get_celebrity_info("1054400")
print(name, name_en, avatar, detail, brief)
'''

if __name__ == '__main__':
    print(get_celebrity_info(1032440))

