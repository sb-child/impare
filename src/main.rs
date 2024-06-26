use impare::libpare_sys::{self, sys::SysApi};

fn main() {
    let req = libpare_sys::sys::CreateTaskRequest {
        data_shards: 1,
        parity_shards: 1,
    };
    let x = libpare_sys::sys::SysApiImpl::create_task(req);
    println!("from go: {}", x.id);
    println!("Hello, world!");
}
