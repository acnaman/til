#! python3
# mcb.pyw  - クリップボードのテキストを保存・復元

import shelve, pyperclip, sys

mcb_shlef = shelve.open('mcb')

# クリップボードの内容を保存
if len (sys.argv) == 3 and sys.argv[1].lower() == 'save':
    mcb_shlef[sys.argv[2]] = pyperclip.paste()
    print('キー\"{}\"に、値\"{}\"を保存しました'.format(sys.argv[2],mcb_shlef[sys.argv[2]]))

elif len(sys.argv) == 2:
    # キーワード一覧と、内容の読み込み
    if sys.argv[1].lower() == 'list':
        pyperclip.copy(str(list(mcb_shlef.keys())))
        print('現在保存されているキーのリスト：{}'.format(str(list(mcb_shlef.keys()))))

    else:
        pyperclip.copy(mcb_shlef[sys.argv[1]])
        print(mcb_shlef[sys.argv[1]])


mcb_shlef.close()
