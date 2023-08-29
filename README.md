# STRUCTURE
```

├── docs // Các tệp OpenAPI/Swagger, tệp định nghĩa giao thức JSON schema.
├── cmd // main
├── config // Chứa cấu hình
│ ├── config.go // Biến môi trường
│ └── constants.go // Biến hằng số
├── docker-compose.yml // Cấu hình container
├── Dockerfile // Định nghĩa img container
├── domain // Lớp miền, giao diện cho lớp cơ sở dữ liệu
│ ├── entity // khai báo đối tượng
│ └── repository // interface của đối tượng
├── go.mod
├── go.sum
├── handler // Xử lý http
├── infra // Lớp cơ sở dữ liệu
│ ├── mysql
│ ├── database.go
│ ├── logger.go
│ ├── model // model cơ sở dữ liệu
│ └── repository // Repository cơ sở dữ liệu
├── logs // Chứa tệp nhật ký
├── Makefile // Các tập lệnh để thực hiện các hoạt động xây dựng, cài đặt, phân tích, v.v.
├── middlewares // Xử lý trước và sau khi xử lý yêu cầu
├── pkg // Các tiện ích cho dịch vụ
├── README.md
├── routers // Bộ định tuyến cho dịch vụ sử dụng REST API
├── scripts
└── usecase // Logic kinh doanh
```

# GIT RULES

## Git Flow Workflow

The Git Flow workflow consists of the following main branches and their purposes:
- `master`: Represents the stable production-ready code. Only merge releases into this branch.
- `develop`: Serves as the integration branch for ongoing development. Feature branches are merged into this branch.
- `feature`: Used to develop new features. Each feature should have its own branch branched off from `develop`.
- `release`: Created for preparing a new release. Bug fixes and last-minute changes can be made in this branch.
- `hotfix`: Created to quickly address critical issues in the production code. Branched off from `master`.

## Branches in Git Flow

In Git Flow, branches have specific naming conventions to indicate their purpose:
- Feature branch: `feature/<branch-name>`
- Release branch: `release/<version-number>`
- Hotfix branch: `hotfix/<branch-name>`

## Using Rebase in Git Flow

To incorporate rebase into the Git Flow workflow, follow these steps:
1. Start a new feature: Create a new feature branch from `develop`.
2. Work on the feature: Make commits to the feature branch as you develop the new feature.
3. Update the feature branch: Before completing the feature, rebase it onto the latest `develop` branch to incorporate the latest changes.
4. Resolve conflicts: If conflicts occur during the rebase process, resolve them by editing the conflicting files manually.
5. Complete the feature: Once the feature is complete and conflicts are resolved, merge the feature branch into `develop`.
6. Repeat the process: Continue working on new features or start a release/hotfix branch following the same principles.
