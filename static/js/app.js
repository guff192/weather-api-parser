let form = Vue.component('current-weather-form', {
    template: '#current-weather-form-template',
    methods: {}
})

let widget = Vue.component('current-weather-widget', {
    props: ['city', 'temp', 'feels_like', 'description'],
    template: '#current-weather-widget-template',
})

let app = new Vue({
    el: '#app',
    data: {
        city: "",
        temp: 0.0,
        feels_like: 0.0,
        description: ""
    },
    methods: {
        getWeather: function () {
            let inputCity = document.getElementById("city_i").value

            let urlParts = document.URL.match(/http.*:\/\/(.*):(\d*)\//i)
            let apiUrl = urlParts[0].replace(urlParts[2], "8080") + "weather/" + inputCity

            const vm = this;
            axios.get(apiUrl)
                .then(function (response) {
                    let data = response.data
                    vm.city = data.name
                    vm.temp = data.main.temp
                    vm.feels_like = data.main.feels_like
                    vm.description = data.weather[0].description
                })
        }
    },
})