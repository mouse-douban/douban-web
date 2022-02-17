import os
import time

if __name__ == '__main__':
    while True:
        time.sleep(60 * 60)
        os.system("kill -31 [pid]")  # todo 自动检测 pid
        pass
