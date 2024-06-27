<template>
    <section class="vh-100" style="background-color: #9A616D;">
        <div class="container py-5 h-100">
            <div class="row d-flex justify-content-center align-items-center h-100">
            <div class="col col-xl-10">
                <div class="card" style="border-radius: 1rem;">
                <div class="row g-0">
                    <div class="col-md-6 col-lg-5 d-none d-md-block">
                    <img src="https://mdbcdn.b-cdn.net/img/Photos/new-templates/bootstrap-login-form/img1.webp"
                        alt="login form" class="img-fluid" style="border-radius: 1rem 0 0 1rem;" />
                    </div>
                    <div class="col-md-6 col-lg-7 d-flex align-items-center">
                    <div class="card-body p-4 p-lg-5 text-black">

                        <form>

                        <h5 class="fw-normal mb-3 pb-3" style="letter-spacing: 1px;">Register your account</h5>

                        <div data-mdb-input-init class="form-outline mb-4">
                            <label class="form-label" for="form2Example17">Phone Number</label>
                            <input type="email" v-model="PhoneNumber" id="form2Example17" class="form-control form-control-lg" />
                        </div>

                        <div data-mdb-input-init class="form-outline mb-4">
                            <label class="form-label" for="form2Example27">Password</label>
                            <input type="password" v-model="Password" id="form2Example27" class="form-control form-control-lg" />
                        </div>

                        <div data-mdb-input-init class="form-outline mb-4">
                            <label class="form-label" for="form2Example271">Repeat Password</label>
                            <input type="password" v-model="RepeatPassword" id="form2Example271" class="form-control form-control-lg" />
                        </div>

                        <div class="pt-1 mb-4">
                            <button data-mdb-button-init data-mdb-ripple-init class="btn btn-dark btn-lg btn-block"  @click="register" type="button">Register</button>
                        </div>

                        <router-link to="/login">
                        <p class="mb-5 pb-lg-2" style="color: #393f81;">Have an account? <a
                            style="color: #393f81;" >Login here</a></p>
                        </router-link>
                        
                        </form>

                    </div>
                    </div>
                </div>
                </div>
            </div>
            </div>
        </div>
        </section>
</template>

<script>
import axios from '../../axios/axiosinstance';

export default{
    data(){
        return {
            PhoneNumber : "",
            Password : "",
            RepeatPassword : ""
        }
    },
    methods:{
        register(){
            if(this.PhoneNumber.length != 10) {
                alert("Invalid Phone number")
                return
            }

            if(this.Password == "") {
                alert("Password Should not be empty")
                return
            }

            if(this.Password != this.RepeatPassword){
                alert("Password and Repeat Password must be same")
                return
            }
            
            var self = this
            axios.post(`signup`,{
                PhoneNumber: parseInt(self.PhoneNumber),
                Password: self.Password
            })
            .then(function (response) {
                if(response.data.Status){
                    alert("Data saved Successfully")
                    self.$router.push({ path: '/login' })
                }
            })
            .catch(function (error) {
                alert(error.response.data.Message);
            })
        },
    }
}
</script>

<style>
</style>