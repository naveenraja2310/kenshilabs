<template>
    <div class="modal">
        <div class="modal-content">
            <form>
                <div class="form-group">
                    <label for="TaskName">Task Name</label>
                    <input type="text" class="form-control" id="TaskName" aria-describedby="emailHelp" v-model="TaskName" placeholder="Name">
                </div>
                <div class="form-group">
                    <label for="TaskAssinedTo">Task Assined To</label>
                    <input type="text" class="form-control" id="TaskAssinedTo" placeholder="Assined To" v-model="TaskAssinedTo">
                </div>
                <div class="form-group">
                    <label for="TaskAssinedBy">Task Assined By</label>
                    <input type="text" class="form-control" id="TaskAssinedBy" placeholder="Assined By" v-model="TaskAssinedBy">
                </div>
                <div class="row justify-content-center my-3">
                    <div class="col-md-auto d-flex justify-content-between">
                        <button type="button" class="btn btn-danger mx-2"  @click="$emit('closeModal')">Cancel</button>
                        <button type="button" class="btn btn-primary mx-2" @click="save" v-if="!ModeUpdate">Save</button>
                        <button type="button" class="btn btn-primary mx-2" @click="update" v-else>Update</button>
                    </div>
                </div>
            </form>
        </div>
    </div>
</template>

<script>
import axios from '../../axios/axiosinstance';
export default {
    data(){
        return {
            ModeUpdate:false,
            TaskID:"",
            TaskName:"",
            TaskAssinedTo : "",
            TaskAssinedBy : ""
        }
    },
    props: ["Data"],
    created(){
        if(this.Data.TaskID){
            this.ModeUpdate = true
        }
        this.TaskID = this.Data.TaskID
        this.TaskName = this.Data.TaskName
        this.TaskAssinedTo = this.Data.TaskAssinedTo
        this.TaskAssinedBy = this.Data.TaskAssinedBy
    },
    methods:{
        save(){
            var self = this
            axios.post(`api/tasks`,{
                TaskName:      self.TaskName,
                TaskAssinedTo: self.TaskAssinedTo,
                TaskAssinedBy: self.TaskAssinedBy,
            })
            .then(function () {
                alert("Data saved Successfully")
                self.$emit('closeModal')
            })
            .catch(function (error) {
                console.log(error);
            })
        },
        update(){
            var self = this
            axios.put(`api/tasks/${self.TaskID}`,{
                TaskName:      self.TaskName,
                TaskAssinedTo: self.TaskAssinedTo,
                TaskAssinedBy: self.TaskAssinedBy,
            })
            .then(function () {
                alert("Data updated Successfully")
                self.$emit('closeModal')
            })
            .catch(function (error) {
                console.log(error);
            })
        }
    }

}
</script>

<style scoped>
    .modal {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    align-items: center;
    }

    .modal-content {
    background-color: #fff;
    padding: 20px;
    border-radius: 5px;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.3);
    max-width: 80%;
    max-height: 80%;
    overflow: auto;
    }

    .modal-item {
    cursor: pointer;
    padding: 10px;
    transition: background-color 0.3s ease;
    }

    .modal-item:hover {
    background-color: #f0f0f0;
    }
</style>