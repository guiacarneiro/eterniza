package controller

//
//func CreateProduto(w http.ResponseWriter, r *http.Request) {
//
//	body, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		util.ERROR(w, http.StatusUnprocessableEntity, err)
//	}
//	user := model.Produto{}
//	err = json.Unmarshal(body, &user)
//	if err != nil {
//		util.ERROR(w, http.StatusUnprocessableEntity, err)
//		return
//	}
//	if err != nil {
//		util.ERROR(w, http.StatusUnprocessableEntity, err)
//		return
//	}
//	userCreated, err := user.Save(database.DB)
//
//	if err != nil {
//
//		formattedError := util.FormatError(err.Error())
//
//		util.ERROR(w, http.StatusInternalServerError, formattedError)
//		return
//	}
//	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
//	util.JSON(w, http.StatusCreated, userCreated)
//}
//
//func GetProdutos(w http.ResponseWriter, r *http.Request) {
//
//	user := model.Produto{}
//
//	users, err := user.FindAllProdutos(database.DB)
//	if err != nil {
//		util.ERROR(w, http.StatusInternalServerError, err)
//		return
//	}
//	util.JSON(w, http.StatusOK, users)
//}
//
//func GetProduto(w http.ResponseWriter, r *http.Request) {
//
//	vars := mux.Vars(r)
//	uid, err := strconv.ParseUint(vars["id"], 10, 32)
//	if err != nil {
//		util.ERROR(w, http.StatusBadRequest, err)
//		return
//	}
//	user := model.Produto{}
//	userGotten, err := user.FindProdutoByID(database.DB, uint32(uid))
//	if err != nil {
//		util.ERROR(w, http.StatusBadRequest, err)
//		return
//	}
//	util.JSON(w, http.StatusOK, userGotten)
//}
//
//func UpdateProduto(w http.ResponseWriter, r *http.Request) {
//
//	vars := mux.Vars(r)
//	uid, err := strconv.ParseUint(vars["id"], 10, 32)
//	if err != nil {
//		util.ERROR(w, http.StatusBadRequest, err)
//		return
//	}
//	body, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		util.ERROR(w, http.StatusUnprocessableEntity, err)
//		return
//	}
//	produto := model.Produto{}
//	produto.ID = uint(uid)
//	err = json.Unmarshal(body, &produto)
//	if err != nil {
//		util.ERROR(w, http.StatusUnprocessableEntity, err)
//		return
//	}
//	tokenID, err := util.ExtractTokenID(r)
//	if err != nil {
//		util.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
//		return
//	}
//	if tokenID != uint32(uid) {
//		util.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
//		return
//	}
//	produto.Prepare()
//	err = produto.Validate("update")
//	if err != nil {
//		util.ERROR(w, http.StatusUnprocessableEntity, err)
//		return
//	}
//	updatedProduto, err := produto.Save(database.DB)
//	if err != nil {
//		formattedError := util.FormatError(err.Error())
//		util.ERROR(w, http.StatusInternalServerError, formattedError)
//		return
//	}
//	util.JSON(w, http.StatusOK, updatedProduto)
//}
//
//func DeleteProduto(w http.ResponseWriter, r *http.Request) {
//
//	vars := mux.Vars(r)
//
//	user := model.Produto{}
//
//	uid, err := strconv.ParseUint(vars["id"], 10, 32)
//	if err != nil {
//		util.ERROR(w, http.StatusBadRequest, err)
//		return
//	}
//	tokenID, err := util.ExtractTokenID(r)
//	if err != nil {
//		util.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
//		return
//	}
//	if tokenID != 0 && tokenID != uint32(uid) {
//		util.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
//		return
//	}
//	_, err = user.DeleteProduto(database.DB, uint32(uid))
//	if err != nil {
//		util.ERROR(w, http.StatusInternalServerError, err)
//		return
//	}
//	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
//	util.JSON(w, http.StatusNoContent, "")
//}
