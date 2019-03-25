# coding: utf-8
require_relative 'yakushoku'

describe Yakushoku do
    let(:yakushoku) {Yakushoku.new}

    example '社員のベース給料は基本給と一緒。基本給が100なら給料も100' do
        expect(yakushoku.calculate_salary(100)).to eq 100
    end

    describe Tanto do
        let(:yakushoku) {Tanto.new}

        example '担当は新入社員なので給料は基本給と一緒。基本給が100なら有料も100' do
            expect(yakushoku.calculate_salary(100)).to eq 100
        end

        example '役職が上がると、主任と同じく基本給の2倍+1' do
            expect(yakushoku.up.calculate_salary(100)).to eq 201
        end
    end

    describe Shunin do
        let(:yakushoku) {Shunin.new}

        example '主任の給料は基本給の2倍+1、基本給が100なら給料は201' do
            expect(yakushoku.calculate_salary(100)).to eq 201
        end

        example '降格したら給料は下がります' do
            expect(yakushoku.down.calculate_salary(100)).to eq 100
        end
    end
end


