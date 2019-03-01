#! python3
# mcb.pyw  - クリップボードのテキストを保存・復元

import shelve pyperclip, sys

mcb_shlef = shelve.open('mcb')

# クリップボードの内容を保存
if len (sys.argv) == 3 and sys.argv[1].lower() == 'save':
    mcb_shlef[sys.argv[2]] = pyperclip.paste()

elif len(sys.argv) == 2:
    # TODO: キーワード一覧と、内容の読み込み


mcb_shlef.close()
