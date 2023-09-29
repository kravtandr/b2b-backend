import xml.etree.ElementTree as ET
import pandas 
from types import NoneType
import os
import requests

def main():
    url = "https://bi-tu-bi.ru/api"
    local_url = "http://127.0.0.1:8080"
    auth = {
        "email": "test@mail.ru", 
        "password": "password123"
    }
    auth_req = requests.post(url+'/user/login', json=auth) 
    print("Auth = ",auth_req.status_code)

    df = pandas.read_excel('novo_products.xlsx', names=['Ид_товара', 'НомерВерсии', 'ПометкаУдаления', 'Штрихкод', 'Артикул', 
                                 'БазоваяЕдиница', 'Ид_категории', "Название_категории","Количество", 'Наименование', 'Описание', "Цена", 
                                'Картинка', 'Страна', 'Вес'])
    # [7:12]
    lineNumber = 0
    for row in df.itertuples():
        if lineNumber==10:
                    break
        # print(row[8:14])
        getCategory_req = requests.post('http://localhost:8080/search/categories?skip=0&limit=1', json={'name': row[8]}) 
        json = getCategory_req.json()
        print("Category id = ", json['data'][0]['id'])
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
           row[13]
        ]
    }
        addProduct_req =requests.post(url+'/product/add', json=product, cookies=auth_req.cookies)
        print(lineNumber," addProduct = ",addProduct_req.status_code)
        lineNumber+=1
    print("Done")


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



if __name__ == "__main__":
     main()