# Ecommerce API - Go Project

Chào mừng bạn đến với dự án **Ecommerce API**. Đây là một hệ thống ecommerce đơn giản được xây dựng bằng ngôn ngữ lập trình Go. Hệ thống hỗ trợ các tính năng như quản lý sản phẩm, giỏ hàng, thanh toán, và quản lý người dùng.

## Nội dung

- [Giới thiệu](#giới-thiệu)
- [Cài đặt](#cài-đặt)
- [Sử dụng](#sử-dụng)
- [API Endpoints](#api-endpoints)
- [Cấu trúc Dự án](#cấu-trúc-dự-án)
- [Công nghệ sử dụng](#công-nghệ-sử-dụng)
- [Kiểm thử](#kiểm-thử)
- [Liên hệ](#liên-hệ)

---

## Giới thiệu

Dự án này cung cấp một API backend cho hệ thống ecommerce, hỗ trợ các chức năng chính như:

- Quản lý người dùng (Đăng ký, Đăng nhập, Xác thực)
- Quản lý sản phẩm (Thêm, Sửa, Xoá)
- Quản lý giỏ hàng (Thêm sản phẩm vào giỏ, Xem giỏ hàng)
- Thanh toán đơn hàng
- Quản lý đơn hàng (Trạng thái, Lịch sử)

## Cài đặt

Để cài đặt và chạy dự án, làm theo các bước dưới đây:

### 1. Cài đặt Go

Trước tiên, bạn cần cài đặt Go. Bạn có thể tải Go tại [https://golang.org/dl/](https://golang.org/dl/).

### 2. Clone Dự Án

Clone repo này về máy của bạn:

```bash
git clone https://github.com/DoanDuong98/go-eccommerce.git
cd go-ecommerce-be
```

### 3. Cài đặt Dependencies

```bash
go mod tidy
```

### 4. Cấu hình môi trường

```bash
DATABASE_URL=your_database_url
JWT_SECRET=your_jwt_secret
PORT=8080
```

### 5. Chạy dự án

```bash
go run main.go
```

## Cấu trúc Dự án

Dự án có cấu trúc thư mục như sau:
```bash
ecommerce-go/
│
├── cmd/                 # Lệnh chính, khởi tạo server
│   └── server.go
│
├── config/              # Cấu hình môi trường
│   └── config.go
│
├── controllers/         # Các controller xử lý request
│   ├── auth_controller.go
│   ├── product_controller.go
│   └── cart_controller.go
│
├── models/              # Các mô hình (models) dữ liệu
│   ├── user.go
│   ├── product.go
│   └── order.go
│
├── services/            # Các logic nghiệp vụ (services)
│   ├── auth_service.go
│   └── product_service.go
│
├── repositories/        # Các thao tác với cơ sở dữ liệu (repositories)
│   ├── user_repo.go
│   ├── product_repo.go
│   └── order_repo.go
│
├── utils/               # Các hàm tiện ích
│   └── jwt.go
│
└── go.mod               # Các phụ thuộc của dự án
```

## Công nghệ sử dụng

- **Go**: Ngôn ngữ chính cho việc xây dựng API backend.
- **GORM**: ORM cho Go để thao tác với cơ sở dữ liệu.
- **JWT**: Để xác thực người dùng.
- **PostgreSQL**: Cơ sở dữ liệu quan hệ (có thể thay thế bằng MySQL hoặc SQLite).
- **Gin**: Web framework cho Go để xây dựng API.
