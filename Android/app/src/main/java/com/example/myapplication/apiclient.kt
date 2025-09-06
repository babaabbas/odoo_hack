package com.example.myapplication

import com.example.myapplication.ui.theme.CreateProjectRequest
import com.example.myapplication.ui.theme.CreateProjectResponse
import com.example.myapplication.ui.theme.CreateTaskRequest
import com.example.myapplication.ui.theme.CreateTaskResponse
import com.example.myapplication.ui.theme.CreateUserRequest
import com.example.myapplication.ui.theme.CreateUserResponse
import com.example.myapplication.ui.theme.ProjectResponse
import retrofit2.http.*
import retrofit2.Retrofit
import retrofit2.converter.gson.GsonConverterFactory

object RetrofitClient {
    private const val BASE_URL = "https://1f37f29a3cbd.ngrok-free.app/"

    private val retrofit: Retrofit by lazy {
        Retrofit.Builder()
            .baseUrl(BASE_URL)
            .addConverterFactory(GsonConverterFactory.create())
            .build()
    }

    val api: ApiService by lazy {
        retrofit.create(ApiService::class.java)
    }
}


interface ApiService {
    @POST("api/users")
    suspend fun createUser(@Body req: CreateUserRequest): CreateUserResponse

    @POST("api/projects")
    suspend fun createProject(@Body req: CreateProjectRequest): CreateProjectResponse

    @GET("api/projects/{id}")
    suspend fun getProjectById(@Path("id") id: Int): ProjectResponse

    @POST("api/tasks")
    suspend fun createTask(@Body req: CreateTaskRequest): CreateTaskResponse
}