pub mod binding {
    rust2go::r2g_include_binding!();
}

#[derive(rust2go::R2G)]
pub struct DemoRequest {
    pub name: String,
    pub age: u8,
}

#[derive(rust2go::R2G)]
pub struct DemoResponse {
    pub pass: bool,
}

#[rust2go::r2g]
pub trait DemoCall {
    fn demo_check(req: DemoRequest) -> DemoResponse;
    fn demo_check_async(req: DemoRequest) -> impl std::future::Future<Output = DemoResponse>;
}
