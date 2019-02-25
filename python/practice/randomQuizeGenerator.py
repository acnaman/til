#! python3
# randomQuizGenerator.py - ランダム順位問題と答えを並べ問題集と回答集を作る

import random #

# 問題のデータ。キーが都道府県で値が県庁所在地
capitals = {
    '北海道': '札幌市', '青森県': '青森市', '岩手県': '盛岡市', '宮城県': '仙台市', '秋田県': '秋田市', '山形県': '山形市', '福島県': '福島市', '茨城県': '水戸市', '栃木県': '宇都宮市', '群馬県': '前橋市', '埼玉県': 'さいたま市', '千葉県': '千葉市', '東京都': '新宿区', '神奈川県': '横浜市', '新潟県': '新潟市', '富山県': '富山市', '石川県': '金沢市', '福井県': '福井市', '山梨県': '甲府市', '長野県': '長野市', '岐阜県': '岐阜市', '静岡県': '静岡市', '愛知県': '名古屋市', '三重県': '津市', '滋賀県': '大津市', '京都府': '京都市', '大阪府': '大阪市', '兵庫県': '神戸市', '奈良県': '奈良市', '和歌山県': '和歌山市', '鳥取県': '鳥取市', '島根県': '松江市', '岡山県': '岡山市', '広島県': '広島市', '山口県': '山口市', '徳島県': '徳島市', '香川県': '高松市', '愛媛県': '松山市', '高知県': '高知市', '福岡県': '福岡市', '佐賀県': '佐賀市', '長崎県': '長崎市', '熊本県': '熊本市', '大分県': '大分市', '宮崎県': '宮崎市', '鹿児島県': '鹿児島市', '沖縄県': '那覇市'
}

numOfChildren = 35

for quiz_num in range(numOfChildren):
    # 問題集と回答集のファイルを作成する

    # 問題集のヘッダーを書く

    # 都道府県の順番をシャッフルする

    # 47都道府県をループして、それぞれ問題を作る



