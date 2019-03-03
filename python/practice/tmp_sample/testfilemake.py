# !python3
# filedelete.pyの動作確認用削除ファイル生成

for x in range(5):
    filename = 'deletefile_{}'.format(x+1)
    with open(filename, 'w', encoding='utf-8') as f:
        f.write('これは削除されるファイル')
