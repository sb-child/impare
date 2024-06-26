use std::{fs, path::Path};

fn main() {
    // path::PathBuf
    fs::remove_file(Path::new("libpare").join("main.go")).unwrap_or(());
    rust2go::Builder::new()
        .with_go_src(Path::new("libpare"))
        .with_regen(
            Path::new("src")
                .join("libpare_sys")
                .join("sys.rs")
                .to_str()
                .unwrap(),
            Path::new("libpare").join("main.go").to_str().unwrap(),
        )
        .build();
    // println!("cargo::rerun-if-changed=src/libpare_sys/sys.rs");
}
