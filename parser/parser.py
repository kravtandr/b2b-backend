import xml.etree.ElementTree as ET
import pandas as pd
from types import NoneType
import os

class Product:
    start = 0
    end = 0
    data = ""
    Id = "null" 
    CategoryId = "null"
    CategoryName = "null"
    Amount = "null"
    Name = "null"
    Info = "null"
    Price = "null"
    Photo = "null"



def ParseGoods(path, files, categories):
    prices = pd.DataFrame(columns=['Id','Price'])
    rests = pd.DataFrame(columns=['Id','Amount'])
    data = pd.DataFrame(columns=['Ид_товара', 'НомерВерсии', 'ПометкаУдаления', 'Штрихкод', 'Артикул', 
                                 'БазоваяЕдиница', 'Ид_категории', "Название_категории","Количество", 'Наименование', 'Описание', "Цена", 
                                'Картинка', 'Страна', 'Вес'])
        # # Чтение Товаров из 1
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
    
    # df = pd.DataFrame(data)
    # df.to_excel('data_'+path+'.xlsx', index=False)
    print("Done import", path)
    return data


def main():
    # Чтение Категорий из корня
    categories = pd.DataFrame(columns=['Id','CategoryName'])
    data = pd.DataFrame(columns=['Ид_товара', 'НомерВерсии', 'ПометкаУдаления', 'Штрихкод', 'Артикул', 
                                 'БазоваяЕдиница', 'Ид_категории', "Название_категории","Количество", 'Наименование', 'Описание', "Цена", 
                                'Картинка', 'Страна', 'Вес'])

    with open('import___af38b3d8-b665-42d4-b0bc-cc696f728ef7.xml', 'r', encoding='utf-8', newline='') as xml_file:
        tree = ET.parse(xml_file)
        root = tree.getroot()
        for category in root.iter('{urn:1C.ru:commerceml_3}Группа'):
            # print(category.tag, category.text)
            Id = category.find('{urn:1C.ru:commerceml_3}Ид').text
            CategoryName= category.find('{urn:1C.ru:commerceml_3}Наименование').text
        
            categories = categories._append({'Id': Id, 'CategoryName': CategoryName}, ignore_index=True)
        categories_df = pd.DataFrame(categories)
        categories_df.to_excel('category.xlsx', index=False)


    # Указываем путь к директории
    basePath = "./goods"
    
    # Получаем список файлов
    files = os.listdir(basePath)
    # print(files)
    files.remove(".DS_Store")
    files.sort()
    for directory in files:
        path = basePath+'/'+directory
        files = os.listdir(basePath+'/'+directory)
        files.remove("import_files")
        files.sort()
        # print(files)
        data = data._append(ParseGoods(path, files, categories))
        # print(os.listdir(basePath+'/'+directory))
    print("All done. save in ./data.xlsx")
    df = pd.DataFrame(data)
    df.to_excel('data.xlsx', index=False)
    # Выводим список файлов
    # print(files)
     


        




if __name__ == "__main__":
     main()



# def main():
#     print("Hello World!")
#     f = open('goods/1/import___97cf0a59-7d9d-4a5e-9dbe-80d54bcdb49f.xml', 'r', encoding='utf-8', newline='')
#     findProducts = False
#     findProduct = False
#     findProductStart = 0
#     findProductEnd = 0
#     lineNumber = 0
#     product = Product()
#     products = []
#     parsedProducts = 0
#     skip =False
#     for line in f:
#         if lineNumber == 500:
#             break
#         line = line.strip()
#         if line.find("<Товары>") != -1:
#             findProducts=True
#         if line.find("</Товары>") != -1:
#             findProducts=False
#             break
#         if findProducts:
#              if line.find("</Товар>") != -1:
#                  findProduct = False
#                  findProductEnd = lineNumber
#                  if product.data != "":
#                     parsedProducts+=1
#                     products.append(product)
#                     product = Product()
#              if line.find("<Товар>") != -1:
#                 findProduct = True
#                 findProductStart = lineNumber
#              if  line.find("<ЗначенияСвойств>") != -1 or line.find("<СтавкиНалогов>") != -1 or line.find("<ЗначенияРеквизитов>") != -1: 
#                 skip = True
#              if findProduct and skip == False:
#                 product.data += line
#              if  line.find("</ЗначенияСвойств>") != -1 or line.find("</СтавкиНалогов>") != -1 or line.find("</ЗначенияРеквизитов>") != -1: 
#                 skip = False

#         print(lineNumber," = ", line)
#         lineNumber+=1
#     f.close()
#     print("End! Parsed Products = ", parsedProducts)
#     print("End! Parsed Products = ", products[0].data)

