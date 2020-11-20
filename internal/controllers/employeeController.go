package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"main/internal/models"
	"main/internal/services"
	"main/internal/utils"
)

type EmployeeController struct {
	EmployeService services.EmployeService
}

func (ctrl *EmployeeController) Save(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.SendErr(w, http.StatusInternalServerError)
	}

	emp := models.Employee{}
	json.Unmarshal(body, &emp)

	errr := ctrl.EmployeService.Save(&emp)
	if errr != nil {
		utils.SendErr(w, http.StatusBadRequest)
	}
	j, _ := json.Marshal(emp)
	utils.SendResponse(w, http.StatusOK, j)
}

func (ctrl *EmployeeController) Get(w http.ResponseWriter, r *http.Request) {

	query, ok := r.URL.Query()["id"]
	if !ok || len(query[0]) < 1 {
		utils.SendErr(w, http.StatusOK)
	}
	id := query[0]
	i, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("%d of type %T", i, i)
	}
	employee, err := ctrl.EmployeService.Get(i)
	if err != nil {
		fmt.Println("Failed to get employtee for id: %s", id)
		utils.SendErr(w, http.StatusInternalServerError)
	}

	j, errr := json.Marshal(employee)
	if errr != nil {
		fmt.Println("Failed to Marshal for id: %s", id)
		utils.SendErr(w, http.StatusInternalServerError)
	}
	utils.SendResponse(w, http.StatusOK, j)
}

func (ctrl *EmployeeController) GetAll(w http.ResponseWriter, r *http.Request) {
	employees, err := ctrl.EmployeService.GetAll()
	if err != nil {
		//fmt.Println("Failed to get employtee for id: %s", id)
		utils.SendErr(w, http.StatusInternalServerError)
	}

	j, _ := json.Marshal(employees)
	utils.SendResponse(w, http.StatusOK, j)
}
