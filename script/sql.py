from __future__ import unicode_literals

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

    def fetch_all(self, sql):
        with self.conn.cursor() as cur:
            cur.execute(sql)
            self.conn.commit()
            return cur.fetchall()

    def execute(self, sql):
        with self.conn.cursor() as cur:
            infect = cur.execute(sql)
            self.conn.commit()
            return infect

    def __fetch_one(self, sql):
        with self.conn.cursor() as cur:
            cur.execute(sql)
            self.conn.commit()
            return cur.fetchone()

    def close(self):
        self.conn.close()

    def insert_subject(self, mid, tags, date, stars, detail, name, score, plot, avatar, celebrities):
        sql = "INSERT INTO subject(mid, tags, date, stars, detail, name, score, plot, avatar, celebrities)" \
              f" VALUES('{mid}', '{tags}', '{date}', '{stars}', '{detail}', '{name}', '{score}', '{plot}', '{avatar}', '{celebrities}')"
        # sql = sql.replace("\"", "")
        self.execute(sql)
        self.conn.commit()

    def insert_celebrity(self, cid, name, name_en, gender, sign, birth, hometown, job, imdb, brief, avatar):
        sql = f"INSERT INTO celebrity(id, name, name_en, gender, sign, birth, hometown, job, imdb, brief, avatar) VALUES('{cid}', '{name}', '{name_en}', '{gender}', '{sign}', '{birth}', '{hometown}', '{job}', '{imdb}', '{brief}', '{avatar}')"
        self.execute(sql)
        self.conn.commit()
