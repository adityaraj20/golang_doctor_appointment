package handler

import (
	"log"
	"net/http"
	userpb "main.go/gunk/v1/user"
)



type UserFilter struct {
	Users []PatientCreate
	SearchTerm string
}


func (h Handler) Show(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatalf("%#v", err)
	}
	st := r.FormValue("SearchTerm")
	ListUser, err := h.usermgmService.UserList(r.Context(),&userpb.UserlistRequest{
		SearchTerm: st,
	})
	if err != nil {
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
	}

	data := []PatientCreate{}
	if ListUser != nil {
		for _, v := range ListUser.GetUsers() {
			data = append(data,PatientCreate{
				ID:        int(v.ID),
				FirstName: v.FirstName,
				LastName:  v.LastName,
				Email:     v.Email,
				Role:      v.Role,
				Username:  v.Username,
				Is_active:    v.IsActive,
			} )
		}
	}	
	Data := UserFilter{
		Users:      data,
		SearchTerm: st,
	}

	h.ParsePatientListTemplate(w, Data)
}

func (h Handler) ParsePatientListTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("listPatient.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}