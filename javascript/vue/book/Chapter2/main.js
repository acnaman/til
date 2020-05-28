var app = new Vue({
  el: '#app',
  data: {
    message: 'Hello, Vue.js!',
    count :0,
    isChild: true,
    isActive: true,
    textColor: 'red',
    bgColor: 'lightgray',
    radius: 0
  },
  methods: {
    increment: function() {
      this.count += 1
    }
  }
})