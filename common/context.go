package common

type Metadata struct {
	Runtime string `yaml:"runtime,omitempty"`
}

type Context struct {
	Meta    Metadata `yaml:"_meta"`
	Name    string   `yaml:"name"`
	Image   string   `yaml:"image"`
	Command []string `yaml:"command"`
	Volumes []string `yaml:"volumes"`
	Detach  bool     `yaml:"detach"`
	TTY     bool     `yaml:"tty"`
	Init    bool     `yaml:"init"`
	Rm      bool     `yaml:"rm"`
}
