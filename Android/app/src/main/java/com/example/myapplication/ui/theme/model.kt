package com.example.myapplication.ui.theme

data class CreateUserRequest(
    val name: String,
    val email: String,
    val password: String
)

data class CreateUserResponse(val id: Int)

data class CreateProjectRequest(
    val name: String,
    val owner_id: Int
)

data class CreateProjectResponse(val id: Int)

data class CreateTaskRequest(
    val project_id: Int,
    val name: String,
    val status: String = "todo",
    val assigned_to: Int? = null
)

data class CreateTaskResponse(val id: Int)

data class ProjectResponse(
    val id: Int,
    val name: String,
    val owner_id: Int,
    val created_at: String,
    val updated_at: String
)