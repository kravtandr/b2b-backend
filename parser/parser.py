import xml.etree.ElementTree as ET
import pandas as pd
from types import NoneType
import os
import re


# files[x] - тут менять индексы, если появились лишние файлы в катологе
def ParseGoods(path, files, categories):
    prices = pd.DataFrame(columns=['Id','Price'])
    rests = pd.DataFrame(columns=['Id','Amount'])
    data = pd.DataFrame(columns=['Ид_товара', 'НомерВерсии', 'ПометкаУдаления', 'Штрихкод', 'Артикул', 
                                 'БазоваяЕдиница', 'Ид_категории', "Название_категории","Количество", 'Наименование', 'Описание', "Цена", 
                                'Картинка', 'Страна', 'Вес'])
    # Чтение Товаров из папок
    with open(path+'/'+files[1], 'r', encoding='utf-8', newline='') as xml_file:
        tree = ET.parse(xml_file)
        root = tree.getroot()
        for price in root.iter('{urn:1C.ru:commerceml_3}Предложение'):
            # print(price.tag, price.text)
            Id = price.find('{urn:1C.ru:commerceml_3}Ид').text
            price= price.find('{urn:1C.ru:commerceml_3}Цены').find('{urn:1C.ru:commerceml_3}Цена').find('{urn:1C.ru:commerceml_3}ЦенаЗаЕдиницу').text       
            prices = prices._append({'Id': Id, 'Price': price}, ignore_index=True)
        prices_df = pd.DataFrame(prices)
        # print(prices.loc[prices['Id'] == '34f0baf5-7924-11e4-8798-10bf48207869'].iat[0, 1])
    print("Done prices", path)
    with open(path+'/'+files[2], 'r', encoding='utf-8', newline='') as xml_file:
        tree = ET.parse(xml_file)
        root = tree.getroot()
        for rest in root.iter('{urn:1C.ru:commerceml_3}Предложение'):
            # print(price.tag, price.text)
            Id = rest.find('{urn:1C.ru:commerceml_3}Ид').text
            rest= rest.find('{urn:1C.ru:commerceml_3}Остатки').find('{urn:1C.ru:commerceml_3}Остаток').find('{urn:1C.ru:commerceml_3}Количество').text       
            rests = rests._append({'Id': Id, 'Amount': rest}, ignore_index=True)
        rests_df = pd.DataFrame(rests)
        # print(rests.loc[rests['Id'] == '34f0baf5-7924-11e4-8798-10bf48207869'].iat[0, 1])
    print("Done rests", path)
    with open(path+'/'+files[0], 'r', encoding='utf-8', newline='') as xml_file:
        tree = ET.parse(xml_file)
        root = tree.getroot()
        # print(ET.tostring(root, encoding='utf8').decode('utf8'))
        for product in root.iter('{urn:1C.ru:commerceml_3}Товар'):
            
            # print(product.tag, product.text)
            Id =product.find('{urn:1C.ru:commerceml_3}Ид').text
            VersionNumber = product.find('{urn:1C.ru:commerceml_3}НомерВерсии').text
            DelMark = product.find('{urn:1C.ru:commerceml_3}ПометкаУдаления').text
            Barcode = product.find('{urn:1C.ru:commerceml_3}Штрихкод').text
            VendorCode = product.find('{urn:1C.ru:commerceml_3}Артикул').text
            Name = product.find('{urn:1C.ru:commerceml_3}Наименование').text
            BasedElem = product.find('{urn:1C.ru:commerceml_3}БазоваяЕдиница').text
            CategoryId = product.find('{urn:1C.ru:commerceml_3}Группы').find('{urn:1C.ru:commerceml_3}Ид').text
            Info = product.find('{urn:1C.ru:commerceml_3}Описание').text
            if type(product.find('{urn:1C.ru:commerceml_3}Картинка')) != NoneType:
                Photo = path+'/'+product.find('{urn:1C.ru:commerceml_3}Картинка').text
            Country = product.find('{urn:1C.ru:commerceml_3}Страна').text
            Weight = product.find('{urn:1C.ru:commerceml_3}Вес').text
            if categories.loc[categories['Id'] == CategoryId].size != 0:
                CategoryName = categories.loc[categories['Id'] == CategoryId].iat[0, 1]
            else:
                CategoryName = "ERROR"
            if prices.loc[prices['Id'] == Id].size != 0:
                Price = prices.loc[prices['Id'] == Id].iat[0, 1]
            else:
                Price = "ERROR"
            if rests.loc[rests['Id'] == Id].size != 0:
                Amount = rests.loc[rests['Id'] == Id].iat[0, 1]
            else:
                Amount = "ERROR"
            data = data._append({'Ид_товара': Id, 'НомерВерсии': VersionNumber, 'ПометкаУдаления': DelMark, 
                                        'Штрихкод': Barcode, 'Артикул': VendorCode, 'Наименование': Name, 
                                        'БазоваяЕдиница': BasedElem, 'Ид_категории': CategoryId, 'Название_категории': CategoryName,'Количество': Amount,  'Описание': Info, 
                                        'Цена': Price, 'Картинка': Photo, 'Страна': Country, 'Вес': Weight}, ignore_index=True)
    
    print("Done import", path)
    return data

def parseCategories(categoriesFilename):
    # Чтение Категорий из корня
    categories = pd.DataFrame(columns=['Id','CategoryName'])

    with open(categoriesFilename, 'r', encoding='utf-8', newline='') as xml_file:
        tree = ET.parse(xml_file)
        root = tree.getroot()
        for category in root.iter('{urn:1C.ru:commerceml_3}Группа'):
            # print(category.tag, category.text)
            Id = category.find('{urn:1C.ru:commerceml_3}Ид').text
            CategoryName= category.find('{urn:1C.ru:commerceml_3}Наименование').text
        
            categories = categories._append({'Id': Id, 'CategoryName': CategoryName}, ignore_index=True)
        categories_df = pd.DataFrame(categories)
        categories_df.to_excel('category.xlsx', index=False)
        return categories

def main():
    data = pd.DataFrame(columns=['Ид_товара', 'НомерВерсии', 'ПометкаУдаления', 'Штрихкод', 'Артикул', 
                                 'БазоваяЕдиница', 'Ид_категории', "Название_категории","Количество", 'Наименование', 'Описание', "Цена", 
                                'Картинка', 'Страна', 'Вес'])

    # находим имя файла с категориями
    files = os.listdir()
    mylist = files 
    r = re.compile(".*\.xml")
    newlist = list(filter(r.match, mylist))
    categories = [] 
    if newlist!=[]:
        print(newlist)
        categoriesFilename = newlist[0]
        # парсим категории
        categories = parseCategories(categoriesFilename)
    else:
        print("No categories files in root")
        return


    # Указываем путь к директории
    basePath = "./goods"

    # Получаем список файлов
    files = os.listdir(basePath)
    print(files)
    try:
        files.remove(".DS_Store")
    except:
        pass

    files.sort()
    for directory in files:
        path = basePath+'/'+directory
        files = os.listdir(basePath+'/'+directory)
        files.remove("import_files")
        files.sort()
        data = data.append(ParseGoods(path, files, categories))
    print("All done. save in ./data.xlsx")
    df = pd.DataFrame(data)
    df.to_excel('data.xlsx', index=False)


if __name__ == "__main__":
     main()
