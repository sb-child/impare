use std::fs;

fn main() {
    fs::remove_file("libpare/main.go").unwrap_or(());
    rust2go::Builder::new()
        .with_go_src("./libpare")
        .with_regen("src/libpare_sys/sys.rs", "libpare/main.go")
        .build();
    // println!("cargo::rerun-if-changed=src/libpare_sys/sys.rs");
}
