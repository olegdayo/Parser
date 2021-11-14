import string
import requests
from bs4 import BeautifulSoup
import time


def get_tags_ebay(ask: str) -> list:
    url = f'https://ru.ebay.com/sch/i.html?_from=R40&_trksid=p2380057.m570.l1313&_nkw={ask}&_sacat=0'
    response = requests.get(url)
    soup = BeautifulSoup(response.text, 'lxml')
    tags = list(map(lambda z: z.text, soup.find('div', class_='srp-related-searches').findAll('span', class_='BOLD')))
    return tags


def get_tags_wildberries(ask: str) -> list:
    url = f'https://www.wildberries.ru/catalog/0/search.aspx?sort=popular&search={ask}/recommendation/catalog?type=maybeyouinterest'
    response = requests.get(url)
    soup = BeautifulSoup(response.text, 'lxml')
    related = soup.findAll('div', id='mainContainer')
    for cell in related:
        print(cell.text)
    print(related)

    # TODO...

    tags = []

    # TODO...

    return tags


def get_tags(ask: str) -> list:
    for symb in list(string.punctuation) + [' ']:
        ask = ask.replace(symb, '+')
    answer = []
    #answer += get_tags_ebay(ask)
    answer = get_tags_wildberries(ask)
    return answer


if __name__ == '__main__':
    start = time.time()
    print(get_tags('аниме'))
    finish = time.time()
    print(str(int((finish - start) * 1000)) + 'ms')
