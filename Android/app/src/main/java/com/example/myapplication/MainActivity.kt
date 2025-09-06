package com.example.myapplication

import ProjectViewModel
import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.activity.enableEdgeToEdge
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.material3.Divider
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import androidx.lifecycle.viewmodel.compose.viewModel
import com.example.myapplication.ui.theme.CreateUserRequest
import com.example.myapplication.ui.theme.MyApplicationTheme
import com.example.myapplication.ui.theme.ProjectResponse

class MainActivity : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        enableEdgeToEdge()
        setContent {
            MyApplicationTheme {
                val viewModel: ProjectViewModel = viewModel() // get ViewModel instance

                ProjectScreen(viewModel = viewModel, projectId = 1)
            }
        }
    }
}

@Composable
fun Greeting(name: String, modifier: Modifier = Modifier) {
    var display by remember { mutableStateOf("Initial Value") }
    LaunchedEffect(Unit) {
        display=cool()
    }

    Text(
        text = "Hello $name! $display",
        modifier = modifier
    )
}
suspend fun cool(): String{
    val user = RetrofitClient.api.createUser(
        CreateUserRequest("John Doe", "john@example.com", "password123")
    )
    val cool="Created user ID: ${user.id}"
    return cool
}
@Composable
fun ProjectScreen(viewModel: ProjectViewModel, projectId: Int) {
    val project by viewModel.project.collectAsState()
    val error by viewModel.error.collectAsState()

    LaunchedEffect(projectId) {
        viewModel.loadProject(projectId)
    }

    Column(modifier = Modifier.padding(16.dp)) {
        error?.let {
            Text(text = "Error: $it", color = Color.Red)
        }

        project?.let { proj ->
            Text("Project: ${proj.name}")
            Text("Owner ID: ${proj.owner_id}")
            Text("Created At: ${proj.created_at}")
        }
    }
}
@Preview(showBackground = true)
@Composable
fun GreetingPreview() {
    MyApplicationTheme {
        Greeting("Android")
    }
}

@Preview
@Composable
fun ProjectScreen2(projectId: Int=1) {
    // Coroutine scope for API calls
    val scope = rememberCoroutineScope()

    // State to hold project data
    var project by remember { mutableStateOf<ProjectResponse?>(null) }
    var error by remember { mutableStateOf<String?>(null) }
    var isLoading by remember { mutableStateOf(true) }

    // Load project once
    LaunchedEffect(projectId) {
        try {
            isLoading = true
            project = RetrofitClient.api.getProjectById(projectId)
        } catch (e: Exception) {
            error = e.message
        } finally {
            isLoading = false
        }
    }

    Column(modifier = Modifier.fillMaxSize().padding(16.dp)) {
        if (isLoading) {
            Text("Loading...")
        } else if (error != null) {
            Text("Error: $error", color = MaterialTheme.colorScheme.error)
        } else if (project != null) {
            Text("Project: ${project!!.name}", style = MaterialTheme.typography.headlineSmall)
            Text("Owner ID: ${project!!.owner_id}")
            Text("Created At: ${project!!.created_at}")

            Spacer(Modifier.height(16.dp))

            Text("Tasks", style = MaterialTheme.typography.titleMedium)
            Spacer(Modifier.height(8.dp))
            Text(project!!.name)

        }
    }
}