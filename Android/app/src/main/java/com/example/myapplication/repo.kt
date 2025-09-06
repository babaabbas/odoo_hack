import com.example.myapplication.ApiService
import com.example.myapplication.ui.theme.ProjectResponse
import retrofit2.HttpException

class ProjectRepository(private val api: ApiService) {

    suspend fun getProjectById(id: Int): Result<ProjectResponse> {
        return try {
            val project = api.getProjectById(id)
            Result.success(project)
        } catch (e: HttpException) {
            val errorBody = e.response()?.errorBody()?.string()
            Result.failure(Exception("HTTP ${e.code()}: $errorBody"))
        } catch (e: Exception) {
            Result.failure(e)
        }
    }
}