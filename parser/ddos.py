import base64
import io
import xml.etree.ElementTree as ET
import pandas 
from types import NoneType
import os
import time
import requests
from PIL import Image
import os
import pandas as pd

def custom_post_request(url, json, cookies):
    max_retries = 2
    retry_delay = 1

    for i in range(max_retries):
        try:
            response = requests.post(url, json=json, cookies=cookies)
            response.raise_for_status()
            return response
        except requests.exceptions.RequestException:
            if i == max_retries - 1:
                raise
            else:
                print("Error occurred. Retrying in", retry_delay, "seconds...")
                time.sleep(retry_delay)

def pillow_image_to_base64_string(img, ext):
    buffered = io.BytesIO()
    img.save(buffered, format="PNG")
    return base64.b64encode(buffered.getvalue()).decode("utf-8")

def base64_string_to_pillow_image(base64_str):
    return Image.open(io.BytesIO(base64.decodebytes(bytes(base64_str, "utf-8"))))


def getProductCategory(product_category):
    categories_df = pandas.read_excel('MappedCategories.xlsx', names=['Id','CategoryName','DBCategory'])
    for row in categories_df.itertuples():  
        if product_category == row[2]:
            print("ProductCategory found = ", row[3])
            return row[3]
    print("ProductCategory not found")
    return "Test"                 

def main():
    url = "https://bi-tu-bi.ru/api"
    local_url = "http://127.0.0.1:8080"
    
    # for local db 
    # url = local_url
    
    print("DDOS:", url)

    #auth
    auth = {
        "email": "novo@test.ru", 
        "password": "Xx2OkpyF"
    }
    auth_req = requests.post(url+'/user/login', json=auth) 
    time.sleep(200/1000)


    if auth_req.status_code == 200:
        
        ddos_log = pd.DataFrame(columns=['name', 'info', 'price', 'docs', 'category_id', 
                                 'amount', 'payWay', "deliveryWay","adress", 'product_photo','code', 'response', 'exception'])
        
        df = pandas.read_excel('data.xlsx', names=['Ид_товара', 'НомерВерсии', 'ПометкаУдаления', 'Штрихкод', 'Артикул', 
                                    'БазоваяЕдиница', 'Ид_категории', "Название_категории","Количество", 'Наименование', 'Описание', "Цена", 
                                    'Картинка', 'Страна', 'Вес'])
        
        # [7:12]
        lineNumber = 0
        err = 0
        stopLine = 8588 #8588
        startLine = 1
        getCategory_resp = {}

        for row in df.itertuples():
            if lineNumber>= startLine and lineNumber<=stopLine:
                product = {}
                Photo = ""
                dataurl = ""
                base64img = ""
                resized_image = Image
                addProduct_req = requests
                # print(row[10][0:15]+"... ", end="")
                Photo = row[13]

                #compress
                compressBy = 2
                ext     = Photo.split('.')[-1]
                image = Image.open(Photo)
                image = image.convert('RGB')
                width, height = image.size
                new_size = (width//compressBy, height//compressBy)
                resized_image = image.resize(new_size)
                base64img=pillow_image_to_base64_string(resized_image, ext)
                dataurl = f'data:image/{ext};base64,{base64img}'
                try: 
                    base64_string_to_pillow_image(base64img)
                except:
                    print("ERROR in photo")
                try:
 
                    db_category_name =getProductCategory(row[8])
                    
                    getCategory_req = requests.post(url + '/search/categories?skip=0&limit=1', json={'name': db_category_name}) 
                    # time.sleep(50/1000)
                    if getCategory_req.status_code != 200 or getCategory_req.json() == {'data': None, 'msg': 'OK'}:
                         print("GetCategory Failed", getCategory_req.status_code, getCategory_req.json())
                    else:
                        getCategory_resp = getCategory_req.json()
                        print("getCategory_resp = ", getCategory_resp)

                        product = {
                            "name": row[10],
                            "info": row[11],
                            "price": int(float(row[12])),
                            "docs":  [],
                            "category_id": getCategory_resp['data'][0]['id'],
                            "amount": int(row[9]),
                            "payWay": "Default",
                            "deliveryWay": "Default",
                            "adress": "Default",
                            "product_photo": [
                            dataurl
                        ]
                        }
                    addProduct_req = custom_post_request(url+'/product/add', json=product, cookies=auth_req.cookies)
                    print(row[10][0:15]+"... ",lineNumber,"/",stopLine," addProduct = ",addProduct_req.status_code, ext)
                    time.sleep(200/1000)
        
                except Exception as e:
                    print(e)
                    err+=1
                    print(row[10][0:15]+"... ",lineNumber,"/",stopLine,"Error addProduct")
                    ddos_log.append({
                        "name": row[10],
                        "info": row[11],
                        "price": float(row[12]),
                        "docs":  [],
                        "category_id": getCategory_resp, # full responce 
                        "amount": float(row[9]),
                        "payWay": "Default",
                        "deliveryWay": "Default",
                        "adress": "Default",
                        "product_photo": [
                        dataurl
                        ],
                        "code": 0,
                        "response": "",
                        "exception": e
                    }, ignore_index=True)
                lineNumber+=1
            else:
                lineNumber+=1
        print("Done. ", "Failed = ",err, "/", lineNumber, " | ", (lineNumber-err)/lineNumber*100, "%")
        print("Save log in ./ddos_log.xlsx")
        df = pd.DataFrame(ddos_log)
        df.to_excel('ddos_log.xlsx', index=False)
        print(ddos_log)
    else:
        print("Auth Failed", auth_req.status_code)
        print("Cred: ", auth)


if __name__ == "__main__":
     main()