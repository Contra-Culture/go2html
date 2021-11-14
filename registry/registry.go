package registry

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Contra-Culture/go2html"
)

type Registry map[string]interface{} // interface{} can contain values of *Template or Registry type

func New() Registry {
	return map[string]interface{}{}
}

func (d Registry) Mkdir(path []string) (Registry, error) {
	if len(path) < 1 {
		return nil, errors.New("wrong path, should not be empty")
	}
	chunk := path[0]
	if len(chunk) < 1 {
		return nil, errors.New("wrong path, path chunk should not be empty")
	}
	i, exists := d[chunk]
	if !exists {
		dir := Registry{}
		d[chunk] = dir
		if len(path) == 1 {
			return dir, nil
		}
		return dir.Mkdir(path[1:])
	}
	switch dir := i.(type) {
	case Registry:
		if len(path) == 1 {
			return dir, nil
		}
		return dir.Mkdir(path[1:])
	default:
		return nil, fmt.Errorf("naming conflict with \"%s\" path: there is already a template with \"%s\" name", strings.Join(path, "/"), path[len(path)-1])
	}
}
func (d Registry) Mkdirf(path []string, f func(dir Registry)) (Registry, error) {
	dir, err := d.Mkdir(path)
	if err != nil {
		return nil, err
	}
	f(dir)
	return dir, nil
}
func (d Registry) Dir(path []string) (Registry, error) {
	if len(path) < 1 {
		return nil, errors.New("wrong path, should not be empty")
	}
	chunk := path[0]
	if len(chunk) < 1 {
		return nil, errors.New("wrong path, path chunk should not be empty")
	}
	i, exists := d[chunk]
	if !exists {
		return nil, fmt.Errorf("wrong path, dir \"%s\" is not found", chunk)
	}
	switch dir := i.(type) {
	case Registry:
		if len(path) == 1 {
			return dir, nil
		}
		return dir.Dir(path[1:])
	default:
		return nil, fmt.Errorf("type error with \"%s\" path: found template with \"%s\" name instead of directory", strings.Join(path, "/"), path[len(path)-1])
	}
}
func (d Registry) Add(t *go2html.Template, path []string) (err error) {
	if len(path) < 1 {
		return errors.New("wrong path, should have at least one chunk")
	}
	if len(path) == 1 {
		_, exists := d[path[0]]
		if exists {
			return fmt.Errorf("template \"%s\" already exists", strings.Join(path, "/"))
		}
		d[path[0]] = t
		return
	}
	prevIdx := len(path) - 1
	dir, err := d.Mkdir(path[:prevIdx])
	if err != nil {
		return
	}
	_, exists := dir[path[prevIdx]]
	if exists {
		return fmt.Errorf("template \"%s\" already exists", strings.Join(path, "/"))
	}
	dir[path[prevIdx]] = t
	return
}
func (d Registry) T(path []string, key string, configure func(*go2html.TemplateCfgr)) (err error) {
	t := go2html.NewTemplate(key, configure)
	if t == nil {
		return fmt.Errorf("template \"%s\" with key \"%s\" somehow was not created", strings.Join(path, "/"), key)
	}
	return d.Add(t, path)
}
func (d Registry) Get(path []string) (t *go2html.Template, err error) {
	if len(path) < 1 {
		return nil, fmt.Errorf("wrong path \"%s\", should have at least one chunk", strings.Join(path, "/"))
	}
	if len(path) == 1 {
		i, exists := d[path[0]]
		if !exists {
			return nil, fmt.Errorf("template \"%s\" does not exist", strings.Join(path, "/"))
		}
		switch t := i.(type) {
		case *go2html.Template:
			return t, nil
		default:
			return nil, fmt.Errorf("type error with \"%s\" path: found directory with \"%s\" name instead of template", strings.Join(path, "/"), path[len(path)-1])
		}
	}
	templateKey := path[len(path)-1]
	dirPath := path[:len(path)-1]
	dir, err := d.Dir(dirPath)
	if err != nil {
		return
	}
	i, exists := dir[templateKey]
	if !exists {
		return nil, fmt.Errorf("template \"%s\" does not exist", strings.Join(path, "/"))
	}
	switch t := i.(type) {
	case *go2html.Template:
		return t, nil
	default:
		return nil, fmt.Errorf("type error with \"%s\" path: found directory with \"%s\" name instead of template", strings.Join(path, "/"), path[len(path)-1])
	}
}
