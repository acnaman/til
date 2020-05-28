var app = new Vue({
  el: '#app',
  data: {
    items: [
      {id: 1, name: "すらいむ", hp: 100},
      {id: 2, name: "ごぶりん", hp: 300},
      {id: 3, name: "りゅうおう", hp: 1000},
    ],
    newname: "",
    nextid: 4
  },
  methods: {
    addMonster: function() {
      this.items.push({id: this.nextid, name: this.newname, hp:500});
      this.nextid++;
      this.newname = "";
    },
    deleteMonster(index) {
      this.items.splice(index, 1);
    },
    attack(index) {
      this.items[index].hp -= 50;
      if (this.items[index].hp <= 0) {
        this.items.splice(index, 1);
      }
    }
  }
})