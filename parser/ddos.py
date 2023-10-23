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

def pillow_image_to_base64_string(img, ext):
    buffered = io.BytesIO()
    img.save(buffered, format="PNG")
    return base64.b64encode(buffered.getvalue()).decode("utf-8")

def base64_string_to_pillow_image(base64_str):
    return Image.open(io.BytesIO(base64.decodebytes(bytes(base64_str, "utf-8"))))

def main():
    url = "https://bi-tu-bi.ru/api"
    local_url = "http://127.0.0.1:8080"
    url = local_url
    auth = {
        "email": "test@mail.ru", 
        "password": "password123"
    }
    auth_req = requests.post(url+'/user/login', json=auth) 
    print("Auth = ",auth_req.status_code)
    time.sleep(200/1000)
    if auth_req.status_code == 200:
        df = pandas.read_excel('data.xlsx', names=['Ид_товара', 'НомерВерсии', 'ПометкаУдаления', 'Штрихкод', 'Артикул', 
                                    'БазоваяЕдиница', 'Ид_категории', "Название_категории","Количество", 'Наименование', 'Описание', "Цена", 
                                    'Картинка', 'Страна', 'Вес'])
        # [7:12]
        lineNumber = 0
        err = 0
        # stopLine = 100
        stopLine = 10000
        startLine = 1
        for row in df.itertuples():
            if lineNumber>= startLine and lineNumber<=stopLine:
                product = {}
                Photo = ""
                dataurl=""
                base64img=""
                resized_image= Image
                addProduct_req = requests
                # if lineNumber>stopLine:
                #             lineNumber-=1
                #             break
            #   print(row[8:14])
                print(row[10][0:15]+"... ", end="")
                Photo = row[13]
                #compress
                ext     = Photo.split('.')[-1]
                image = Image.open(Photo)
                image = image.convert('RGB')
                width, height = image.size
                new_size = (width//2, height//2)
                resized_image = image.resize(new_size)
                base64img=pillow_image_to_base64_string(resized_image, ext)
                dataurl = f'data:image/{ext};base64,{base64img}'
                try: 
                    base64_string_to_pillow_image(base64img)
                except:
                    print("ERROR in photo")
                # binary_fc       = open(Photo, 'rb').read()  
                # base64_utf8_str = base64.b64encode(binary_fc).decode('utf-8')
                # photoPath = Photo
                # ext     = photoPath.split('.')[-1]
                # dataurl = f'data:image/{ext};base64,{base64_utf8_str}'
                # Photo = dataurl
                try:
                    getCategory_req = requests.post(local_url + '/search/categories?skip=0&limit=1', json={'name': row[8]}) 
                    time.sleep(100/1000)
                    json = getCategory_req.json()
                    # print("Category id = ", json['data'][0]['id'])
                    product = {
                    "name": row[10],
                    "info": row[11],
                    "price": int(row[12]),
                    "docs":  [],
                    "category_id": json['data'][0]['id'],
                    "amount": int(row[9]),
                    "payWay": "Default",
                    "deliveryWay": "Default",
                    "adress": "Default",
                    "product_photo": [
                    dataurl
                    ]
                }

                    addProduct_req =requests.post(url+'/product/add', json=product, cookies=auth_req.cookies, timeout=5.0)
                    # print("start sleep")
                    time.sleep(1200/1000)
                    # print("stop sleep")
                    print(lineNumber,"/",stopLine," addProduct = ",addProduct_req.status_code, ext)
                    
                except:
                    err+=1
                    print(lineNumber,"/",stopLine,"Error addProduct")
                lineNumber+=1
            else:
                lineNumber+=1
        print("Done. ", "Failed = ",err, "/", lineNumber, " | ", (lineNumber-err)/lineNumber*100, "%")


        # print([['Наименование','Описание', "Цена", 'Название_категории', 'Количество','Картинка',]])

        # print(data_df.head())
        # with open('novo_products.xlsx', 'r', encoding='utf-8') as data:
        #      lineNumber = 0
        #      for line in data:
        #           if lineNumber==10:
        #                break
        #           print(line)
        #           lineNumber+=1


        # r = requests.post('https://bi-tu-bi.ru/api/user/login', json=auth) 
        # print("Auth = ",r)
        # print("Auth = ",r.status_code)
    else:
        print("Auth Failed")
        print("Cred: ", auth)


if __name__ == "__main__":
     main()