# coding: utf-8
require_relative 'shain'

describe Shain do
    let(:shain) {Shain.new}

    example '社員はとりあえず起立します' do
        expect(shain.work).to eq '社員はとりあえず起立する。'
    end

    example '担当は慌てて起立します' do
        shain.gyoumu = TantoGyoumu.new
        expect(shain.work).to eq '担当は慌てて起立しました。'
    end

    example '主任は素早く起立します' do
        shain.gyoumu = ShuninGyoumu.new
        expect(shain.work).to eq '主任が素早く立ちました。'
    end

    example '社員のベース給料は基本給と一緒。基本給が100なら給料も100' do 
        expect(shain.calculate_salary(100)).to eq 100
    end

    context '役職が担当の場合' do
        before { shain.yakushoku = Tanto.new }

        example '役職が担当だと、給料は基本給と一緒' do 
            expect(shain.calculate_salary(100)).to eq 100
        end

        example '昇進すると給料が主任と同じになる' do
            expect(shain.up.calculate_salary(100)).to eq 201
        end
    end

    context '役職が主任の場合' do
        before { shain.yakushoku = Shunin.new }

        example '役職が担当だと、給料は基本給の2倍+1' do 
            expect(shain.calculate_salary(100)).to eq 201
        end

        example '降格すると給料が担当と同じになる' do
            expect(shain.down.calculate_salary(100)).to eq 100
        end
    end
end