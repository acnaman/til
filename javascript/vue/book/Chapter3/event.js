var app = new Vue({
  el: '#app',
  data: {
    show: false,
  },
  methods: {
    handleClick() {
      var count = this.$refs.count;
      if(count) {
        count.innerText = parseInt(count.innerText, 10) + 1
      }
    }
  }
})