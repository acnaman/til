# coding: utf-8
require_relative 'shain'

# Yakushoku class
class Yakushoku
    def calculate_salary(kihonkyu)
        kihonkyu
    end

    def up
    end

    def down
    end
end

class Tanto < Yakushoku
    def up
        Shunin.new
    end
end

class Shunin < Yakushoku
    def calculate_salary(kihonkyu)
        kihonkyu*2 + 1
    end
    
    def down
        Tanto.new
    end
end
