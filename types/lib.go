package types

import (
	"time"
)

//region AUTH

type SecUser struct {
	Model
	Username string      `json:"username"`
	Password string      `json:"password"`
	Email    string      `json:"email"`
	Groups   []*SecGroup `json:"groups" gorm:"many2many:sec_user_groups;"`
}

type SecGroup struct {
	Model
	Code string `json:"code"`
	Name string `json:"name"`
	//Users []*SecUser `json:"users" gorm:"many2many:sec_user_groups;"`
	Perms []*SecPerm `json:"perms" gorm:"many2many:sec_group_perms;"`
}

type SecPerm struct {
	Model
	Code  string `json:"code"`
	Name  string `json:"name"`
	Value string `json:"value"`
	//Groups []*SecGroup `json:"groups" gorm:"many2many:group_perms;"`
}

type Session struct {
	Model
	Username  string
	SessionID string
	Perms     string
}

//endregion AUTH

type Feature struct {
	Model
	Name  string `json:"name" yaml:"name"`
	Desc  string `json:"desc" yaml:"desc"`
	Repo  uint   `json:"repo"`
	Alias string `json:"alias"`
}

type Robot struct {
	Feature
}

type App struct {
	Feature
	IsNative *bool
}

type Service struct {
	Feature
	Ipcid    string `json:"ipcid"`
	Env      string `json:"env" yaml:"env"`
	Enabled  *bool  `yaml:"enabled" json:"enabled"`
	Image    string `yaml:"image" json:"image"`
	Delay    int    `yaml:"delay" json:"delay"`
	Pid      int    `json:"pid" yaml:"-"`
	Hostname string `json:"hostname" yaml:"hostname"`
	Gui      bool   `json:"gui" yaml:"gui"`
}

type Runtime struct {
	Feature
	Path string `json:"path" yaml:"path"`
}

type Repo struct {
	Model
	Feature  string `json:"feature" yaml:"feature"`
	Name     string `json:"name" yaml:"name"`
	Url      string `json:"repo"`
	Alias    string `json:"alias"`
	Cred     uint   `json:"cred"`
	LocalDir string `json:"local_dir"`
	Status   string `json:"status"`
}

type ReplayDescriptor struct {
	Feature         string            `json:"feature" yaml:"feature"`
	ReplayID        string            `json:"replay_id" yaml:"replay_id"`
	Name            string            `json:"name" yaml:"name"`
	Desc            string            `json:"desc" yaml:"desc"`
	Runner          string            `json:"runner" yaml:"runner"`
	Main            string            `json:"main" yaml:"main"`
	Args            []string          `json:"args" yaml:"args"`
	Env             map[string]string `json:"env" yaml:"env"`
	InstallScript   string            `json:"install_script" yaml:"install_script"`
	UninstallScript string            `json:"uninstall_script" yaml:"uninstall_script"`
	AbortScript     string            `json:"abort_script" yaml:"abort_script"`
	Id              string            `json:"id" yaml:"-"`
	LocalDir        string            `json:"local_dir" yaml:"-"`
	RepoID          uint              `json:"repo_id" yaml:"-"`
	QueueID         uint              `json:"queue_id" yaml:"-"`
	Config          map[string]string `json:"config" yaml:"config"`
	Alias           string            `json:"alias" yaml:"alias"`
	Autostart       bool              `json:"autostart" yaml:"autostart"`
	Hostname        string            `json:"hostname" yaml:"hostname"`
	Ipcid           string            `json:"ipcid"`
	Gui             bool              `json:"gui" yaml:"gui"`
	Link            string            `json:"link" yaml:"link"` //This prop allows redirect of repos
}

type NotifyRequest struct {
	Title string
	Msg   string
}

type Model struct {
	ID uint `gorm:"primarykey" json:"id"`
}

//type Model simpledb.SimplePersistent

type Menu struct {
	Model
	Name    string `json:"name"`
	Label   string `json:"label"`
	Tooltip string `json:"tooltip"`
	Target  string `json:"target"`
	Index   int    `json:"index" gorm:"column:menuindex"`
	Enabled *bool  `json:"enabled"`
}

type Config struct {
	Model
	K    string `json:"k"`
	V    string `json:"v"`
	T    string `json:"t"`
	Repo uint   `json:"repo"`
}

type Credentials struct {
	Model
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
	Sshkey   string `json:"sshkey"`
}

type Log struct {
	Model
	Log []byte
}

type Job struct {
	Model
	Name  string `json:"name"`
	Code  string `json:"code"`
	File  string `json:"file"`
	Notes string `json:"notes"`
}

type Route struct {
	Model
	From        string   `json:"from"`
	Host        string   `json:"host"`
	To          string   `json:"to"`
	Desc        string   `json:"desc"`
	Enabled     *bool    `json:"enabled"`
	Index       string   `json:"index"`
	HeadersJson string   `json:"headers_json"`
	Permissions string   `json:"permissions"`
	PermArr     []string `json:"-" gorm:"-"`
}

type Queue struct {
	Model
	Job         uint      `json:"job"`
	Ready       bool      `json:"ready"`
	Processed   bool      `json:"processed"`
	ProcessedAt time.Time `json:"processed_at"`
	Request     []byte
}

//type QueueData struct {
//	Model
//	Qid     uint
//	Request []byte
//}

type CryptoKey struct {
	Model
	Defkey bool
	Name   string
	Pub    []byte
	Priv   []byte
}

type Options struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

//type Version struct {
//	Model
//	Job   string `json:"job"`
//	Ver   string `json:"ver"`
//	Notes string `json:"notes"`
//}
//
//type VersionData struct {
//	Model
//	Vid  uint   `json:"vid"`
//	Data []byte `json:"data"`
//}

type Cron struct {
	Model
	Cron    string `json:"cron"`
	JobCode uint   `json:"job"`
	Enabled *bool  `json:"enabled"`
}

//type Resource struct {
//	Model
//	Name string `json:"name"`
//	Mime string `json:"mime"`
//	Len  int    `json:"len"`
//}
//
//type ResourceContent struct {
//	Model
//	Resid   uint   `json:"resid"`
//	Content []byte `json:"content"`
//}

type License struct {
	Model
	Owner     string `json:"owner"`
	RobotID   string `json:"robot_id"`
	MachineID string `json:"machine_id"`
	UID       string `json:"uid"`
	Alias     string `json:"alias"`
}

type KV struct {
	ID string `gorm:"primarykey" json:"id"`
	V  string `json:"v"`
}

/*
@API
*/
type VerCheckRequest struct {
	Chan string            `json:"chan"`
	Os   string            `json:"os"`
	Arch string            `json:"arch"`
	Data map[string]string `json:"data"`
}

/*
@API
*/
type VerCheckResponse struct {
	Data map[string]string `json:"data"`
}

/*
@API
*/
type VerGetRequest struct {
	Os   string `json:"os"`
	Chan string `json:"chan"`
	Arch string `json:"arch"`
	Id   string `json:"id"`
}

/*
@API
*/
type VerGetResponse struct {
	Data []byte `json:"data"`
}

type SQLRequest struct {
	Sql string `json:"sql"`
}
type SQLResponse struct {
	Data interface{} `json:"data"`
	Err  string      `json:"err"`
}

type SecUserChOwnPassRequest struct {
	Oldpass string `json:"oldpass"`
	Newpass string `json:"newpass"`
}

type LoginRequest struct {
	Username string
	Password string
}

type RecipeItem struct {
	Url      string `json:"url" yaml:"url"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	Sshkey   string `json:"sshkey" yaml:"sshkey"`
}

type Recipe struct {
	User     string       `json:"user" yaml:"user"`
	Password string       `json:"password" yaml:"password"`
	Sshkey   string       `json:"sshkey" yaml:"sshkey"`
	LicUser  string       `json:"licuser" yaml:"licuser"`
	License  string       `json:"license" yaml:"license"`
	Alias    string       `json:"alias" yaml:"alias"`
	Items    []RecipeItem `json:"items" yaml:"items"`
}

type PubSubMsg struct {
	Msg  string
	Data interface{}
}

type PubSubMsgRunnerFinishData struct {
	Err  error
	Desc *ReplayDescriptor
}

type Series struct {
	Serie string    `json:"serie"`
	When  time.Time `json:"when"`
	Value float64   `json:"value"`
	Sec   int       `json:"sec"`
	Min   int       `json:"min"`
	Hour  int       `json:"hour"`
	Day   int       `json:"day"`
	Month int       `json:"month"`
	Year  int       `json:"year"`
}

func (s *Series) SetTime(t time.Time) {
	s.When = t
	s.Day = t.Day()
	s.Hour = t.Hour()
	s.Min = t.Minute()
	s.Sec = t.Second()
	s.Month = int(t.Month())
	s.Year = t.Year()
}
