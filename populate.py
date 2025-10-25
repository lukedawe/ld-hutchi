import json
import requests
from dotenv import load_dotenv
import os

def main():
    load_dotenv('api/dev.env')
    port = os.getenv('PORT')
    if(port == None): return 1
    with open('common/dogs.json') as f:
        d = json.load(f)
        # print(d)
        request = {"categories":[]}
        for category, breed in d.items():
            category_entry = {
                "name": category,
                "breeds": [{"name": breed_name} for breed_name in breed]
            }
            request["categories"].append(category_entry)
        request = json.dumps(request, indent=2)
        # print(request)
        response: requests.Response = requests.post('http://localhost:'+port+'/v1/categories', data=request)
        print(response.json())

main()