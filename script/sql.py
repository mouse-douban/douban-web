import json
import pymysql


def load_config():
    with open('../config/config.json') as c:
        return json.load(c)


class DB:
    config = load_config()

    param = {
        "host": config["default_ip_and_port"].split(":")[0],
        "port": int(config["default_ip_and_port"].split(":")[1]),
        "db": config["default_db_name"],
        "user": config["default_root"],
        "password": config["default_password"],
        "charset": config["default_charset"]
    }

    conn = pymysql.connect(**param)

    def __fetch_all(self, sql):
        with self.conn.cursor() as cur:
            cur.execute(sql)
            return cur.fetchall()

    def execute(self, sql):
        with self.conn.cursor() as cur:
            return cur.execute(sql)

    def __fetch_one(self, sql):
        with self.conn.cursor() as cur:
            cur.execute(sql)
            return cur.fetchone()

    def close(self):
        self.conn.close()

    def insert_subject(self, tags, date, stars, detail, name, score, plot, avatar, celebrities):
        self.execute(f"INSERT INTO subject(tags, date, stars, detail, name, score, plot, avatar, celebrities)"
                     f" VALUES('{tags}', '{date}', '{stars}', '{detail}', '{name}', '{score}', '{plot}', '{avatar}', '{celebrities}')")

    def insert_celebrity(self, id, name, name_en, gender, sign, birth, hometown, job, imdb, brief):
        sql = f"INSERT INTO celebrity(id, name, name_en, gender, sign, birth, hometown, job, imdb, brief) VALUES('{id}', '{name}', '{name_en}', '{gender}', '{sign}', '{birth}', '{hometown}', '{job}', '{imdb}', '{brief}')"
        self.execute(sql)


dB = DB()

dB.insert_celebrity(2, "哈哈", "haha", "女", "天鱼座", "2021-02-11 12:00:00", "银河", "马飞飞", "8018hgdaybwd", "yidauw")

dB.close()