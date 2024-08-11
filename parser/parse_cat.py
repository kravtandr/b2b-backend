import time
import requests
import xml.etree.ElementTree as ET
import pandas as pd
from types import NoneType
import os

'''
Баня и отопление
Двери. Окна. Лестницы
Изоляционные материалы
Инструмент
Кровля. Фасад. Забор
Метиз и крепеж
Отделочные материалы
Пиломатериалы
Подарочные карты
Сантехника. Водоснабжение. Канализация
Строительные материалы
Товары для дома и сада
Электротовары
'''
# не трогать - воняет
def main():
    # Чтение Категорий из корня
    categories_subcategories = pd.DataFrame(columns=['Category','SubCategory'])
    data = pd.DataFrame(columns=['Ид_товара', 'НомерВерсии', 'ПометкаУдаления', 'Штрихкод', 'Артикул', 
                                 'БазоваяЕдиница', 'Ид_категории', "Название_категории","Количество", 'Наименование', 'Описание', "Цена", 
                                'Картинка', 'Страна', 'Вес'])
    with open('import___af38b3d8-b665-42d4-b0bc-cc696f728ef7.xml', 'r', encoding='utf-8', newline='') as xml_file:
        tree = ET.parse(xml_file)
        root = tree.getroot()
        print(len(root[0][2][0][4]))
        print(root[0][2][0][4][0][3].text)
        i = 0
        for high_group in root[0][2][0][4]:
            print(i, high_group[3].text)
            i+=1
            if len(high_group)>4:
                subcategories = []
                for category in high_group[4].iter():
                    if (type(category.find('{urn:1C.ru:commerceml_3}Наименование'))!= NoneType):
                        url = "https://bi-tu-bi.ru/api"
                        search_cat = category.find('{urn:1C.ru:commerceml_3}Наименование').text
                        getCategory_req = requests.post(url + '/search/categories?skip=0&limit=1', json={'name': search_cat}) 
                        time.sleep(125/1000)
                        json = getCategory_req.json()
                        time.sleep(25/1000)
                        try:
                            if json['data']==None:
                                if len(search_cat) >= 15:
                                    search_cat = search_cat[:-9]
                                    print(search_cat)
                                    getCategory_req = requests.post(url + '/search/categories?skip=0&limit=1', json={'name': search_cat}) 
                                    time.sleep(125/1000)
                                    json = getCategory_req.json()
                                    time.sleep(25/1000)
                            print("------ ",category.find('{urn:1C.ru:commerceml_3}Наименование').text, "Id = ", json['data'][0]['id'])
                            subcategories.append( json['data'][0]['id'])
                        except:
                            print("------ skip ", category.find('{urn:1C.ru:commerceml_3}Наименование').text)
                categories_subcategories = categories_subcategories._append({'Category':  high_group[3].text, 'SubCategory': subcategories}, ignore_index=True)
    categories_df = pd.DataFrame(categories_subcategories)
    categories_df.to_excel('clastered_cat.xlsx', index=False)


if __name__ == "__main__":
     main()
