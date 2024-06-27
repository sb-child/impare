pub mod binding {
    #![allow(warnings)]
    rust2go::r2g_include_binding!();
}

#[derive(rust2go::R2G)]
#[repr(C)]
/// Create a new task
/// <https://pkg.go.dev/github.com/klauspost/reedsolomon@v1.12.1#New>
pub struct CreateTaskRequest {
    pub data_shards: usize,
    pub parity_shards: usize,
}

#[derive(rust2go::R2G)]
#[repr(C)]
pub struct CreateTaskResponse {
    pub id: u64,
}

#[derive(rust2go::R2G)]
#[repr(C)]
/// Remove a task
pub struct RemoveTaskRequest {
    pub id: u64,
}

#[derive(rust2go::R2G)]
#[repr(C)]
pub struct RemoveTaskResponse {
    /// `true` if the task is successfully removed
    pub success: bool,
}

#[rust2go::r2g(binding)]
#[allow(unused)]
pub trait SysApi {
    fn create_task(req: CreateTaskRequest) -> CreateTaskResponse;
    fn remove_task(req: RemoveTaskRequest) -> RemoveTaskResponse;
}
