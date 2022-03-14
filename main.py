import string
import sys
import time

import requests
from bs4 import BeautifulSoup


# Ebay.
def get_tags_ebay(ask: str) -> list:
    # Creating url using given request.
    url = f'https://ru.ebay.com/sch/i.html?_from=R40&_trksid=p2380057.m570.l1313&_nkw={ask}&_sacat=0'
    # Creating request.
    response = requests.get(url)
    # Creating page source .html document.
    soup = BeautifulSoup(response.text, 'lxml')
    # Searching all needed div tags and all needed span tags.
    tags = list(map(lambda z: z.text, soup.find('div', class_='srp-related-searches').findAll('span', class_='BOLD')))
    # Returning tags.
    return tags


def get_tags_wildberries(ask: str) -> list:
    # Creating url using given request.
    url = f'{ask}'
    # Creating request.
    response = requests.get(url)
    # Creating page source .html document.
    soup = BeautifulSoup(response.text, 'lxml')
    # related = soup.findAll('div', id='mainContainer')
    # for cell in related:
    #     print(cell.text)
    # print(related)

    # TODO...

    # Tags.
    tags = []

    # TODO...

    # Returning tags.
    return tags


# Getting all tags.
def get_tags(ask: str) -> list:
    for symb in list(string.punctuation) + [' ']:
        ask = ask.replace(symb, '+')
    answer = []
    answer += get_tags_ebay(ask)
    answer = get_tags_wildberries(ask)
    return answer


if __name__ == '__main__':
    # Setting timer.
    start = time.time()

    # If string is given, using it, otherwise set "anime", because anime rules.
    if len(sys.argv) > 1:
        print(get_tags(sys.argv[1]))
    else:
        print(get_tags('anime'))

    # Getting end time.
    finish = time.time()
    # Printing the total time elapsed.
    print(str(int((finish - start) * 1000)) + 'ms')
