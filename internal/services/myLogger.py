import logging
logger = logging.getLogger(__name__)
logger.setLevel(logging.INFO)
handler = logging.StreamHandler()
handler.setFormatter(logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s'))
logger.addHandler(handler)

def logger_decorator(func):
    def wrapper(*args, **kwargs):
        results = func(*args, **kwargs)
        logger.info("Вызов функции: %s \n Аргументы: %s, %s \n Результат: %s", func.__name__, args, kwargs, results)
        return results
    return wrapper