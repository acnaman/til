import shelve

def makeshelf():
    shelf_file = shelve.open('mydata')
    cats = ['Zophie', 'Pooka', 'Simon']

    shelf_file['cats'] = cats

    shelf_file.close()

def showshelf():
    shelf_file = shelve.open('mydata')
    print(shelf_file['cats'])
    shelf_file.close()
