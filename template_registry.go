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
		content  map[string]interface{}
	}
)

func Reg(name string) *TemplateRegistry {
	return &TemplateRegistry{
		name: name,
		dirs: map[string]*TemplateRegistryDir{},
	}
}
func (r *TemplateRegistry) Mkdir(path []string) (*TemplateRegistryDir, error) {
	if len(path) < 1 {
		return nil, errors.New("wrong path, should not be empty")
	}
	chunk := path[0]
	if len(chunk) < 1 {
		return nil, errors.New("wrong path, path chunk should not be empty")
	}
	dir, exists := r.dirs[chunk]
	if !exists {
		dir = &TemplateRegistryDir{
			name:     chunk,
			content:  map[string]interface{}{},
		}
		r.dirs[chunk] = dir
	}
	if len(path) == 1 {
		return dir, nil
	}
	return dir.Mkdir(path[1:])
}
func (r *TemplateRegistry) Dir(path []string) (*TemplateRegistryDir, error) {
	if len(path) < 1 {
		return nil, errors.New("wrong path, should not be empty")
	}
	chunk := path[0]
	if len(chunk) < 1 {
		return nil, errors.New("wrong path, path chunk should not be empty")
	}
	dir, exists := r.dirs[chunk]
	if !exists {
		return nil, fmt.Errorf("wrong path, dir \"%s\" is not found", chunk)
	}
	if len(path) == 1 {
		return dir, nil
	}
	return dir.Dir(path[1:])
}
func (r *TemplateRegistry) Add(t *Template, path []string) (err error) {
	if len(path) < 2 {
		return fmt.Errorf("wrong path \"%s\", should have at least two chunks", strings.Join(path, "/"))
	}
	prevIdx := len(path) - 1
	dir, err := r.Mkdir(path[:prevIdx])
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
func (r *TemplateRegistry) Get(path []string) (t *Template, err error) {
	if len(path) < 2 {
		return nil, fmt.Errorf("wrong path \"%s\", should have at least two chunks", strings.Join(path, "/"))
	}
	templateKey := path[len(path)-1]
	dirPath := path[:len(path)-1]
	d, err := r.Dir(dirPath)
	if err != nil {
		return
	}
	i, exists := d.content[templateKey]
	if !exists {
		return nil, fmt.Errorf("template \"%s\" does not exist", strings.Join(path, "/"))
	}
	t, ok := i.(*Template)
	if !ok {
		return nil, fmt.Errorf("type error with \"%s\" path: found directory with \"%s\" name instead of template", strings.Join(path, "/"), path[len(path) - 1])
	}
	return
}

func (d *TemplateRegistryDir) Mkdir(path []string) (*TemplateRegistryDir, error) {
	if len(path) < 1 {
		return nil, errors.New("wrong path, should not be empty")
	}
	chunk := path[0]
	if len(chunk) < 1 {
		return nil, errors.New("wrong path, path chunk should not be empty")
	}
	i, exists := d.content[chunk]
	if !exists {
		dir := &TemplateRegistryDir{
			name:     chunk,
			content:  map[string]interface{}{},
		}
		d.content[chunk] = dir
		if len(path) == 1 {
			return dir, nil
		}
		return dir.Mkdir(path[1:])
	} else {
		dir, ok := i.(*TemplateRegistryDir)
		if !ok {
			return nil, fmt.Errorf("naming conflict with \"%s\" path: there is already a template with \"%s\" name", strings.Join(path, "/"), path[len(path) - 1])
		}
		if len(path) == 1 {
			return dir, nil
		}
		return dir.Mkdir(path[1:])
	}
}
func (d *TemplateRegistryDir) Dir(path []string) (*TemplateRegistryDir, error) {
	if len(path) < 1 {
		return nil, errors.New("wrong path, should not be empty")
	}
	chunk := path[0]
	if len(chunk) < 1 {
		return nil, errors.New("wrong path, path chunk should not be empty")
	}
	i, exists := d.content[chunk]
	if !exists {
		return nil, fmt.Errorf("wrong path, dir \"%s\" is not found", chunk)
	}
	dir, ok := i.(*TemplateRegistryDir)
	if !ok {
		return nil, fmt.Errorf("type error with \"%s\" path: found template with \"%s\" name instead of directory", strings.Join(path, "/"), path[len(path) - 1])
	}
	if len(path) == 1 {
		return dir, nil
	}
	return dir.Dir(path[1:])
}
func (d *TemplateRegistryDir) Add(t *Template, path []string) (err error) {
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
	dir, err := d.Mkdir(path[:prevIdx])
	if err != nil {
		return
	}
	_, exists := dir.content[path[prevIdx]]
	if exists {
		// FIXME: Looks like this if condition is not correct.
		return fmt.Errorf("template \"%s\" does not exist", strings.Join(path, "/"))
	}
	dir.content[path[prevIdx]] = t
	return
}
func (d *TemplateRegistryDir) Get(path []string) (t *Template, err error) {
	if len(path) < 1 {
		return nil, fmt.Errorf("wrong path \"%s\", should have at least one chunk", strings.Join(path, "/"))
	}
	if len(path) == 1 {
		i, exists := d.content[path[0]]
		if !exists {
			return nil, fmt.Errorf("template \"%s\" does not exist", strings.Join(path, "/"))
		}
		t, ok := i.(*Template)
		if !ok {
			return nil, fmt.Errorf("type error with \"%s\" path: found directory with \"%s\" name instead of template", strings.Join(path, "/"), path[len(path) - 1])
		}
		return t, nil
	}
	templateKey := path[len(path)-1]
	dirPath := path[:len(path)-1]
	dir, err := d.Dir(dirPath)
	if err != nil {
		return
	}
	i, exists := dir.content[templateKey]
	if !exists {
		return nil, fmt.Errorf("template \"%s\" does not exist", strings.Join(path, "/"))
	}
	t, ok := i.(*Template)
	if !ok {
		return nil, fmt.Errorf("type error with \"%s\" path: found directory with \"%s\" name instead of template", strings.Join(path, "/"), path[len(path) - 1])
	}
	return
}
