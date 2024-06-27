<template>
    <navbar />
    <div class="cointainer">
        <div class="row">
        <button type="button" class="btn btn-primary" style="margin-bottom: 10px;" @click="AddTask">Add Task</button>
            <table  class="table table-striped table-sm table-bordered">
                <thead>
                    <th>Task Name</th>
                    <th>Task Assigned By</th>
                    <th>Task Assined To</th>
                    <th>Created At</th>
                    <th>Updated At</th>
                    <th>Actions</th>
                </thead>
                <tbody>
                    <tr v-for="task in TaskData" :key="task.TaskID">
                        <td>{{ task.TaskName }} </td>
                        <td>{{ task.TaskAssinedBy }}</td>
                        <td>{{ task.TaskAssinedTo }}</td>
                        <td>{{ formatDate(task.CreatedAt) }}</td>
                        <td>{{ formatDate(task.UpdatedAt) }}</td>
                        <td>
                            <button style="margin-right: 5px;"
                                type="button"
                                class="btn btn-success btn-sm"
                                @click="Edit(task)">
                                <i class="fa fa-pencil" style="color: black;"></i>
                            </button>
                            <button style="margin-right: 5px;"
                                type="button"
                                class="btn btn-danger btn-sm"
                                @click="Delete(task)">
                                <i class="fa fa-trash" style="color: black;"aria-hidden="true"></i>
                            </button>                            
                        </td>
                    </tr>
                </tbody>
            </table>
        </div> 
    </div>
    <Modal v-if="showModal"
        :Data="Data"
        @closeModal="closeModal"
    />
</template>

<script>
import axios from '../../axios/axiosinstance';
import Modal from './taskmodal.vue'
import navbar from './navbar.vue';

export default{
    data(){
        return {
            TaskData:[],
            Data:{},
            showModal:false
        }
    },
    created(){
        this.GetData()
    },
    components:{
        navbar,
        Modal
    },
    methods:{
        GetData(){
            var self = this
            axios.get(`api/tasks`)
            .then(function (response) {
                self.TaskData = response.data.Data
                console.log(response);
            })
            .catch(function (error) {
                console.log(error);
            })
        },
        formatDate(dateString) {
            const zeroDate = "0001-01-01T00:00:00Z";
            if (!dateString || dateString === zeroDate) return '';
            const date = new Date(dateString);
            const day = String(date.getDate()).padStart(2, '0');
            const month = String(date.getMonth() + 1).padStart(2, '0'); 
            const year = date.getFullYear();
            return `${day}/${month}/${year}`;
        },
        AddTask(){
            this.showModal = true
        },
        closeModal(){
            this.showModal = false
            this.Data = {}
            this.GetData()
        },
        Edit(task){
            this.showModal = true,
            this.Data = task
        },
        Delete(task){
            var self = this
            axios.delete(`api/tasks/${task.TaskID}`)
            .then(function () {
                alert("Data deleted Successfully")
                self.GetData()
            })
            .catch(function (error) {
                console.log(error);
            })
        }

    }

}
</script>

<style>
.cointainer{
    margin: 40px;
}
</style>