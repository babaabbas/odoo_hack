import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.example.myapplication.ui.theme.ProjectResponse
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.launch

class ProjectViewModel(private val repository: ProjectRepository) : ViewModel() {

    private val _project = MutableStateFlow<ProjectResponse?>(null)
    val project: StateFlow<ProjectResponse?> = _project

    private val _error = MutableStateFlow<String?>(null)
    val error: StateFlow<String?> = _error

    fun loadProject(id: Int) {
        viewModelScope.launch {
            val result = repository.getProjectById(id)
            result.onSuccess { proj ->
                _project.value = proj
            }.onFailure { err ->
                _error.value = err.message
            }
        }
    }
}