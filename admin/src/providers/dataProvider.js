const axios = require('axios')

const client = axios.create({
    baseURL: 'http://localhost:8080/',
})

const dataProvider = {
    getList: (page, query) => {
        return client.get(page, query).then((res) => {
            console.log(res)

            res.data = res.data.map((x) => {
                x.id = x.uid
                delete x.uid
                return x
            })

            return {
                data: res.data,
                total: res.data.length
            }
        })
    },
    getOne: (resource, params) => {
        return client.get(resource, params).then(res => ({
            data: res.data
        }))
    },
    getMany: (resource, params) => {
        console.log('getMany->' + resource)
        return client.get(resource)
    },
    getManyReference: (resource, params) => {
        console.log('getManyReference->' + resource)
        return client.get(resource)
    },
    update: (resource, params) => {
        
    },
    updateMany: (resource, params) => {

    },
    create: (resource, params) => {

    },
    delete: (resource, params) => {

    },
    deleteMany: (resource, params) => {
        
    }
}

export default dataProvider