## Analisis Entity

- User
- Campaigns
- Campaign Images
- Transactions

---

## Entity Relationship Diagram (ERD)

kita bisa menggunaka erdplus.com

1. Diagram

<img src="./ERD.png" style="align:denter;">

2. Details
   - Users
     - bisa membuat banyak campaign dan bersifat opsional
     - bisa memiliki banyak transaksi dan bersifat opsional
   - Campaign
     - wajib dimiliki oleh 1 user
     - wajib memiliki beberapa gambar campaign
     - bisa memiliki banyak transaksi
   - Campaign Image
     - wajib dimiliki oleh 1 campaign
   - Transactions
     - wajib dimiliki oleh 1 user
     - wajib memiliki oleh 1 campaign

---

## Entity Fields / Columns

1. User

   - id : int
   - name : varchar
   - occupations : varchar
   - email : varchar
   - password_hash : varchar
   - avatar_file_name : varchar
   - role : varchar
   - token : varchar
   - created_at : datetime
   - updated_at : datetime

2. Campaigns

   - id : int
   - user_id : int
   - name : varchar
   - short_description : varchar
   - goal_amount : int
   - current_amount : int
   - description : text
   - perks : text
   - backer_count : int
   - slug : varchar
   - created_at : datetime
   - updated_at : datetime

3. Campaign Images

   - id : int
   - campaign_id : int
   - file_name : varchar
   - is_primary : boolean (tinyint)
   - created_at : datetime
   - updated_at : datetime

4. Transaction
   - id : int
   - campaign_id : int
   - user_id : int
   - amount : int
   - status : varchar
   - code : varchar
   - payment_url: varchar
   - token : varchar
   - created_at : datetime
   - updated_at : datetime

## JWT Auth

jwt dimanfaatkan untuk authentikasi API berdasarkan token user. [Details jwt](http://jwt.io)

```bash
# get jwt module
go get github.com/dgrijalva/jwt-go
```

1. create auth service auth/service.go

   ```go

   // service.go
   package auth
   import "github.com/dgrijalva/jwt-go"

   /**
   1. generate token
   2. validasi token
   */

   type Service interface {
      GenerateToken(userID int) (string, error)
   }

   type jwtService struct {
   }

   func NewService() *jwtService {
      return &jwtService{}
   }

   var SECRET_KEY = []byte("asdhbvsad364t834iqwb")

   func (s *jwtService) GenerateToken(userID int) (string, error) {
      claim := jwt.MapClaims{}
      claim["user_id"] = userID

      // generate token
      token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

      // tanda tangani token
      signedToken, err := token.SignedString(SECRET_KEY)
      if err != nil {
         return signedToken, err
      }

      return signedToken, nil
   }

   ```

2. API response with token

   ```json
   {
     "meta": {
       "message": "Suuccessfully loggedin",
       "code": 200,
       "status": "success"
     },
     "data": {
       "id": 1,
       "name": "test simpan dari service",
       "occupation": "erendhoheiri",
       "email": "contoh@gmail.com",
       "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxfQ.pR5mUyz1tm_Ni6-mCi-ankpmIwVifpJ0k_tNjbyp6p8",
       "ImageUrl": ""
     }
   }
   ```

3. Auth Middleware
   - ambil nilai header Authorization: Bearer token
   - dari Authorization: Bearer, hanya diambil nilai token saja
   - validasi token
   - token => user_id
   - ambil user dari db berdasarkan user_id melalui service
   - set context dangan isi user

## API Contract

**POST : api/v1/users**
**params :**

- name
- occupation
- email
- password

**response :**

```json
meta : {
	message: 'Your account has been created',
	code: 200,
	status: 'success'
},
data : {
	id: 1,
	name: "Agung Setiawan",
	occupation: "content creator",
	email: "com.agungsetiawan@gmail.com",
	token: "peterpanyangterdalam"
}
```

**POST : api/v1/email_checkers
params :**

- email

**response:**

```json
meta : {
	message: 'Email address has been registered',
	code: 200,
	status: 'success'
},
data : {
	is_available: false
}
```

**POST: api/v1/avatars
params:**

- avatar (form)

**response:**

```json
meta : {
	message: 'Avatar successfully uploaded,
	code: 200,
	status: 'success'
},
data : {
	is_uploaded: true
}
```

**POST: api/v1/sessions
params:**

- email
- password

**response:**

meta : {

message: 'You're now logged in'

code: 200

status: 'success'

},

data : {

id: 1,
name: "Agung Setiawan",
occupation: "content creator",

email: "com.agungsetiawan@gmail.com",

token: "peterpanyangterdalam"

}

**GET: api/v1/campaigns
params:**
optional

- user_id
- backer_id
- none

**response:**

```json
meta : {
	message: 'List of campaigns',
	code: 200,
	status: 'success'
},
data : [
{
		id: 1,
		name: "BWA Startup",
		short_description: "Laris manis tanjung kimpul, mari belajar bareng",
		image_url: "domain/path/image.jpg",
		goal_amount: 1000000000,
		current_amount: 500000000,
		slug: "slug-here",
		user_id: 10
	}
]
```

**GET: api/v1/campaigns/1
params:**
none

**response:**

```json
meta : {
	message: 'single campaigns',
	code: 200,
	status: 'success'
},
data : {
	id: 1,
	name: "BWA Startup",
	short_description: "Laris manis tanjung kimpul, mari belajar bareng",
	image_url: "path/image.jpg",
	goal_amount: 1000000000,
	current_amount: 500000000,
	user_id: 10,
  slug: "slug",
	description: "Lorem epsum dolor sit amet yang panjang text-nya",
	user : {
		name: "Julia Ester",
		image_url: "path/image.jpg"
	},
	perks: [
		"Nintendo Switch",
		"Play Station 4"
	],
	images: [
		{
			image_url: "path/image.jpg",
			is_primary: true
		}
	]
}
```

POST: api/v1/campaigns

```json
{
  "meta": {
    "message": "Campaign successfully created ",
    "code": 200,
    "status": "success"
  },
  "data": {
    "user_id": 1,
    "goal_amount": 100000000,
    "current_amount": 0,
    "id": 7,
    "name": "Switch Pro",
    "image_url": "",
    "slug": "slug",
    "short_description": "Upcoming Nintendo Switch Pro"
  }
}
```

PUT : api/v1/campaigns/1

sama dengan atas

POST: api/v1/campaign-images

- file
- campaign_id
- is_primary

sama dengan upload avatar

GET : api/v1/campaigns/:id/transactions (campaign punya transaksi backer siapa aja)

```json
meta : {
	message: 'List of transactions'
	code: 200,
	status: 'success'
},
data : [
	{
		id: 1,
		name: "Agung Setiawan",
		amount: 1000000000,
		created_at: datetime
	}
]
```

GET : api/v1/transactions/ (user pernah transaksi apa aja)
params : header auth (current user)

```json
meta : {
	message: 'List of backed campaigns',
	code: 200,
	status: 'success'

},
data : [
	{
		id: 1,
		amount: 1000000000,
		status: "paid",
		created_at: datetime,
		campaign: {
			name: "Hola",
			image_url : "path/to/file.png"
		}
	}
]
```

POST: api/v1/transactions
params : header auth (current user)

***request***

```json
{
    "campaign_id" : 2,
    "amount" : 200000
}
```


***response***

```json
{
    "meta": {
        "message": "Success to create transaction",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 5,
        "campaign_id": 2,
        "user_id": 2,
        "amount": 200000,
        "status": "pending",
        "code": "",
        "token": "4*****-0655-**********",
        "payment_url": "https://app.sandbox.midtrans.com/snap/v2/vtweb/**************************",
        "created_at": "0001-01-01T00:00:00Z"
    }
}
```