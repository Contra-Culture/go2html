package go2html

import (
	"errors"
	"fmt"
	"strings"
)

type (
	TemplateRegistry struct {
		name string
		dirs map[string]*TemplateRegistryDir
	}
	TemplateRegistryDir struct {
		name     string
		content  map[string]*Template
		children map[string]*TemplateRegistryDir
	}
)

func Reg(name string) *TemplateRegistry {
	return &TemplateRegistry{
		name: name,
		dirs: map[string]*TemplateRegistryDir{},
	}
}
func (r *TemplateRegistry) Mkdir(path ...string) (*TemplateRegistryDir, error) {
	if len(path) < 1 {
		return nil, errors.New("wrong path, should not be empty")
	}
	chunk := path[0]
	if len(chunk) < 1 {
		return nil, errors.New("wrong path, path chunk should not be empty")
	}
	dir, exists := r.dirs[chunk]
	if exists {
		if len(path) == 1 {
			return nil, fmt.Errorf("directory \"%s\" already exists", chunk)
		}
	} else {
		dir = &TemplateRegistryDir{
			name:     chunk,
			content:  map[string]*Template{},
			children: map[string]*TemplateRegistryDir{},
		}
		r.dirs[chunk] = dir
	}
	if len(path) == 1 {
		return dir, nil
	}
	return dir.Mkdir(path[1:]...)
}
func (r *TemplateRegistry) Dir(path ...string) (*TemplateRegistryDir, error) {
	if len(path) < 1 {
		return nil, errors.New("wrong path, should not be empty")
	}
	chunk := path[0]
	if len(chunk) < 1 {
		return nil, errors.New("wrong path, path chunk should not be empty")
	}
	dir, ok := r.dirs[chunk]
	if !ok {
		return nil, fmt.Errorf("wrong path, dir \"%s\" is not found", chunk)
	}
	if len(path) == 1 {
		return dir, nil
	}
	dir, err := dir.Mkdir(path[1:]...)
	if err != nil {
		return nil, err
	}
	return dir, nil
}
func (r *TemplateRegistry) Add(t *Template, path ...string) (err error) {
	if len(path) < 2 {
		return fmt.Errorf("wrong path \"%s\", should have at least two chunks", strings.Join(path, "/"))
	}
	prevIdx := len(path) - 1
	dir, err := r.Mkdir(path[:prevIdx]...)
	if err != nil {
		return
	}
	_, exists := dir.content[path[prevIdx]]
	if exists {
		return fmt.Errorf("template \"%s\" already exists", strings.Join(path, "/"))
	}
	dir.content[path[prevIdx]] = t
	return
}
func (r *TemplateRegistry) Get(path ...string) (t *Template, err error) {
	if len(path) < 2 {
		return nil, fmt.Errorf("wrong path \"%s\", should have at least two chunks", strings.Join(path, "/"))
	}
	templateKey := path[len(path)-1]
	dirPath := path[:len(path)-1]
	d, err := r.Dir(dirPath...)
	if err != nil {
		return
	}
	t, ok := d.content[templateKey]
	if !ok {
		return nil, fmt.Errorf("template \"%s\" does not exist", strings.Join(path, "/"))
	}
	return
}

func (d *TemplateRegistryDir) Mkdir(path ...string) (*TemplateRegistryDir, error) {
	if len(path) < 1 {
		return nil, errors.New("wrong path, should not be empty")
	}
	chunk := path[0]
	if len(chunk) < 1 {
		return nil, errors.New("wrong path, path chunk should not be empty")
	}
	dir, exists := d.children[chunk]
	if exists {
		if len(path) == 1 {
			return nil, fmt.Errorf("directory \"%s\" already exists", chunk)
		}
	} else {
		dir = &TemplateRegistryDir{
			name:     chunk,
			content:  map[string]*Template{},
			children: map[string]*TemplateRegistryDir{},
		}
		d.children[chunk] = dir
	}
	if len(path) == 1 {
		return dir, nil
	}
	return dir.Mkdir(path[1:]...)
}
func (d *TemplateRegistryDir) Dir(path ...string) (*TemplateRegistryDir, error) {
	if len(path) < 1 {
		return nil, errors.New("wrong path, should not be empty")
	}
	chunk := path[0]
	if len(chunk) < 1 {
		return nil, errors.New("wrong path, path chunk should not be empty")
	}
	dir, ok := d.children[chunk]
	if !ok {
		return nil, fmt.Errorf("wrong path, dir \"%s\" is not found", chunk)
	}
	if len(path) == 1 {
		return dir, nil
	}
	dir, err := dir.Mkdir(path[1:]...)
	if err != nil {
		return nil, err
	}
	return dir, nil
}
func (d *TemplateRegistryDir) Add(t *Template, path ...string) (err error) {
	if len(path) < 1 {
		return errors.New("wrong path, should have at least one chunk")
	}
	if len(path) == 1 {
		_, exists := d.content[path[0]]
		if exists {
			return fmt.Errorf("template \"%s\" already exists", path[0])
		}
		d.content[path[0]] = t
		return
	}
	prevIdx := len(path) - 1
	dir, err := d.Mkdir(path[:prevIdx]...)
	if err != nil {
		return
	}
	_, exists := dir.content[path[prevIdx]]
	if exists {
		return fmt.Errorf("template \"%s\" does not exist", strings.Join(path, "/"))
	}
	dir.content[path[prevIdx]] = t
	return
}
func (d *TemplateRegistryDir) Get(path ...string) (t *Template, err error) {
	if len(path) < 1 {
		return nil, fmt.Errorf("wrong path \"%s\", should have at least one chunk", strings.Join(path, "/"))
	}
	if len(path) == 1 {
		t, ok := d.content[path[0]]
		if !ok {
			return nil, fmt.Errorf("template \"%s\" does not exist", strings.Join(path, "/"))
		}
		return t, nil
	}
	templateKey := path[len(path)-1]
	dirPath := path[:len(path)-1]
	dir, err := d.Dir(dirPath...)
	if err != nil {
		return
	}
	t, ok := dir.content[templateKey]
	if !ok {
		return nil, fmt.Errorf("template \"%s\" does not exist", strings.Join(path, "/"))
	}
	return
}
