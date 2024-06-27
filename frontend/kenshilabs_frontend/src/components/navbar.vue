<template>
    <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
    <router-link to="/">
        <a class="navbar-brand">Kenshi Labs</a>
    </router-link>
    
    <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
    </button>

    <div class="collapse navbar-collapse" id="navbarSupportedContent">
        <ul class="navbar-nav mr-auto">
        </ul>
        <form class="form-inline my-2 my-lg-0">
            <router-link to="/login">
                <button class="btn btn-outline-success my-2 my-sm-0" style="margin-right:10px" type="button" >Login</button>
                <button class="btn btn-outline-danger my-2 my-sm-0" type="button" @click="Logout">Log Out</button>
			</router-link>        
        </form>
    </div>
    </nav>
</template>

<script>

import axios from '../../axios/axiosinstance'
export default {
    data(){
        return{

        }
    },
    methods:{
        Logout(){
            var self = this
            axios.post(`signout`)
            .then(function (res) {
                console.log(res);
            })
            .catch(function (error) {
                console.log(error);
            })
        }
    }
}
</script>

