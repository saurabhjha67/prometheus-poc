import json
import os
import argparse

if __name__ == '__main__':
    parser = argparse.ArgumentParser()

    parser.add_argument("--project", required=True)

    parser.add_argument("--MSG_TMP_DIR", required=True)

    parser.add_argument("--MSG_BODY_FILE_NAME", required=True)

    args = parser.parse_args()

    msg = {  "messages": [{"attributes": {"test": "test" },"data": {"project_id": args.project}} ]}

    if not os.path.exists(args.MSG_TMP_DIR):
        os.makedirs(f"{args.MSG_TMP_DIR}/")

    with open(f"{args.MSG_TMP_DIR}/{args.MSG_BODY_FILE_NAME}", "w") as fp:
        json.dump(msg, fp)
