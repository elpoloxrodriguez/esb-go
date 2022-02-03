(window["webpackJsonp"] = window["webpackJsonp"] || []).push([["registrarse-registrarse-module"],{

/***/ "./node_modules/raw-loader/index.js!./src/app/registrarse/registrarse.page.html":
/*!*****************************************************************************!*\
  !*** ./node_modules/raw-loader!./src/app/registrarse/registrarse.page.html ***!
  \*****************************************************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = "<nb-layout center>\n\n    \n    <nb-layout-column center>\n\n\n        <nb-card accent=\"success\" [nbSpinner]=\"loading\"\n        nbSpinnerStatus=\"danger\"\n             nbSpinnerSize=\"large\">\n       <nb-card-header >\n\n\n            <h1 id=\"title\" class=\"title\">Registrarse</h1>\n            <p class=\"sub-title\">Hola! Bienvenidos a Ipsfa en línea.</p>\n   \n        </nb-card-header>\n        <nb-card-body><br>\n          <div class=\"row\">\n                <div class=\"col-lg-12\">\n\n                        \n                        <form (ngSubmit)=\"add('top-right', 'danger', dialog)\" #form=\"ngForm\" aria-labelledby=\"title\">\n\n                            <div class=\"form-control-group\">\n                                \n                                <label class=\"label\" for=\"input-tipo\">Tipo:</label>\n                                <nb-select placeholder=\"Seleccioné tipo de afiliado\" \n                                  fullWidth fieldSize=\"large\" [(ngModel)]=\"titular\" \n                                  name=\"titular\" (selectedChange)=\"seleccionarTipo($event)\" >\n                                    <nb-option value=\"TIT\">Titular</nb-option>\n                                    <nb-option value=\"SOB\">Sobreviviente</nb-option>\n                                    <!-- <nb-option value=\"SOB-TIT\">Titular - Sobreviviente</nb-option> -->\n                                </nb-select>\n                            </div>\n\n                                <div class=\"form-control-group\">\n                                \n                                  <label class=\"label\" for=\"input-email\">{{ lblcedula }}:</label>\n                                  <input nbInput\n                                         fullWidth\n                                         name=\"usuario\"\n                                         id=\"input-cedula\"\n                                         placeholder=\"{{ lblcedula }}\"\n                                         fieldSize=\"large\"\n                                         [(ngModel)]=\"usuario\" \n                                         value=\"{{usuario}}\"\n                                         autofocus>\n                                </div>\n                                \n\n                                \n\n                                <div class=\"form-control-group\">\n                                \n                                    <label class=\"label\" for=\"input-email\">Componente:</label>\n                                    <nb-select placeholder=\"Seleccioné componente de procedencia\" \n                                    [(ngModel)]=\"componente\" name=\"componente\" fullWidth fieldSize=\"large\">\n                                        <nb-option value=\"EJ\">Ejercito Bolivariano</nb-option>\n                                        <nb-option value=\"AR\">Armada Bolivariana</nb-option>\n                                        <nb-option value=\"AV\">Aviación Bolivariana</nb-option>\n                                        <nb-option value=\"GN\">Guardia Nacional Bolivariana</nb-option>\n                                    </nb-select>\n                                </div>\n\n                               <div class=\"form-control-group\">\n                                \n                                    <label class=\"label\" for=\"input-email\">Correo Electrónico:</label>\n                                    <input nbInput\n                                           fullWidth\n                                           name=\"correo\"\n                                           id=\"input-correo\"\n                                           placeholder=\"Correo Electrónico\"\n                                           fieldSize=\"large\"\n                                           [(ngModel)]=\"correo\" \n                                           value=\"{{correo}}\"\n                                           autofocus>\n                                </div>\n\n                                <div class=\"form-control-group\">\n                                  <span class=\"label-with-link\">\n                                    <label class=\"label\" for=\"input-password\">Clave:</label>\n                                    \n                                  </span>\n                                  <input nbInput\n                                         fullWidth\n                                         name=\"clave\"\n                                         type=\"password\"\n                                         id=\"input-password\"\n                                         placeholder=\"Clave\"\n                                         fieldSize=\"large\"\n                                         [(ngModel)]=\"clave\" \n                                         value=\"{{clave}}\">\n                                </div>\n                                <div class=\"form-control-group\">\n                                    <span class=\"label-with-link\">\n                                      <label class=\"label\" for=\"input-password\">Repetir Clave:</label>\n                                    </span>\n                                    <input nbInput\n                                           fullWidth\n                                           name=\"rclave\"\n                                           type=\"password\"\n                                           id=\"input-rpassword\"\n                                           placeholder=\"Repetir Clave\"\n                                           fieldSize=\"large\"\n                                           [(ngModel)]=\"rclave\" \n                                           value=\"{{rclave}}\"\n                                           >\n                                  </div>\n                                \n                                <br><br>\n                                <div class=\"row\">\n                                    <div class=\"col-lg-4\">\n                                        <button nbButton\n                                        fullWidth\n                                        status=\"\"\n                                        size=\"medium\"\n                                        routerLink=\"../login\">\n                                        Atrás\n                                      </button>\n\n                                    </div>\n                                    <div class=\"col-lg-4\"></div>\n                                    <div class=\"col-lg-4\"><button nbButton\n                                        fullWidth\n                                        status=\"success\"\n                                        size=\"medium\"\n                                        [disabled]=\"submitted || !form.valid\"\n                                        [class.btn-pulse]=\"submitted\"\n                                        [nbSpinner]=\"loading\" nbSpinnerStatus=\"info\" \n                                        nbSpinnerSize=\"medium\" nbSpinnerMessage=\"\">\n                                  Registrate\n                                </button></div>\n                                    \n                                </div>\n                                \n                                \n                              </form>\n                </div>\n        </div>\n      </nb-card-body>\n        \n      </nb-card>\n\n    </nb-layout-column>\n    \n\n\n</nb-layout>\n\n\n<ng-template #dialog let-data let-ref=\"dialogRef\">\n  <nb-card>\n    <nb-card-header>Mensaje Ipsfa en linea</nb-card-header>\n    <nb-card-body>\n      <br> Se ha registrado correctamente el afiliado. \n      <br> Bienvenidos a ipsfa en linea\n    </nb-card-body>\n    <nb-card-footer>\n      <button nbButton (click)=\"continuar()\" status-success>Continuar</button>\n    </nb-card-footer>\n  </nb-card>\n</ng-template>\n\n"

/***/ }),

/***/ "./src/app/registrarse/registrarse.module.ts":
/*!***************************************************!*\
  !*** ./src/app/registrarse/registrarse.module.ts ***!
  \***************************************************/
/*! exports provided: RegistrarsePageModule */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "RegistrarsePageModule", function() { return RegistrarsePageModule; });
/* harmony import */ var tslib__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! tslib */ "./node_modules/tslib/tslib.es6.js");
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/core */ "./node_modules/@angular/core/fesm2015/core.js");
/* harmony import */ var _angular_common__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! @angular/common */ "./node_modules/@angular/common/fesm2015/common.js");
/* harmony import */ var _angular_forms__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! @angular/forms */ "./node_modules/@angular/forms/fesm2015/forms.js");
/* harmony import */ var _angular_router__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! @angular/router */ "./node_modules/@angular/router/fesm2015/router.js");
/* harmony import */ var _nebular_theme__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! @nebular/theme */ "./node_modules/@nebular/theme/fesm2015/index.js");
/* harmony import */ var _theme_theme_module__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! ../@theme/theme.module */ "./src/app/@theme/theme.module.ts");
/* harmony import */ var _registrarse_page__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! ./registrarse.page */ "./src/app/registrarse/registrarse.page.ts");




 // we also need angular router for Nebular to function properly



const routes = [
    {
        path: '',
        component: _registrarse_page__WEBPACK_IMPORTED_MODULE_7__["RegistrarsePage"],
    }
];
let RegistrarsePageModule = class RegistrarsePageModule {
};
RegistrarsePageModule = tslib__WEBPACK_IMPORTED_MODULE_0__["__decorate"]([
    Object(_angular_core__WEBPACK_IMPORTED_MODULE_1__["NgModule"])({
        imports: [
            _theme_theme_module__WEBPACK_IMPORTED_MODULE_6__["ThemeModule"],
            _angular_common__WEBPACK_IMPORTED_MODULE_2__["CommonModule"],
            _nebular_theme__WEBPACK_IMPORTED_MODULE_5__["NbActionsModule"],
            _nebular_theme__WEBPACK_IMPORTED_MODULE_5__["NbCheckboxModule"],
            _nebular_theme__WEBPACK_IMPORTED_MODULE_5__["NbDatepickerModule"], _nebular_theme__WEBPACK_IMPORTED_MODULE_5__["NbIconModule"],
            _nebular_theme__WEBPACK_IMPORTED_MODULE_5__["NbInputModule"],
            _nebular_theme__WEBPACK_IMPORTED_MODULE_5__["NbRadioModule"],
            _nebular_theme__WEBPACK_IMPORTED_MODULE_5__["NbSelectModule"],
            _nebular_theme__WEBPACK_IMPORTED_MODULE_5__["NbSpinnerModule"],
            _nebular_theme__WEBPACK_IMPORTED_MODULE_5__["NbTabsetModule"],
            _nebular_theme__WEBPACK_IMPORTED_MODULE_5__["NbListModule"],
            _nebular_theme__WEBPACK_IMPORTED_MODULE_5__["NbUserModule"],
            _angular_forms__WEBPACK_IMPORTED_MODULE_3__["FormsModule"],
            _angular_router__WEBPACK_IMPORTED_MODULE_4__["RouterModule"].forChild(routes),
            _angular_router__WEBPACK_IMPORTED_MODULE_4__["RouterModule"],
            _nebular_theme__WEBPACK_IMPORTED_MODULE_5__["NbLayoutModule"],
            _nebular_theme__WEBPACK_IMPORTED_MODULE_5__["NbSidebarModule"],
            _nebular_theme__WEBPACK_IMPORTED_MODULE_5__["NbButtonModule"],
            _nebular_theme__WEBPACK_IMPORTED_MODULE_5__["NbAlertModule"],
            _nebular_theme__WEBPACK_IMPORTED_MODULE_5__["NbCardModule"],
        ],
        declarations: [_registrarse_page__WEBPACK_IMPORTED_MODULE_7__["RegistrarsePage"]]
    })
], RegistrarsePageModule);



/***/ }),

/***/ "./src/app/registrarse/registrarse.page.scss":
/*!***************************************************!*\
  !*** ./src/app/registrarse/registrarse.page.scss ***!
  \***************************************************/
/*! no static exports found */
/***/ (function(module, exports) {

module.exports = "\n/*# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6IiIsImZpbGUiOiJzcmMvYXBwL3JlZ2lzdHJhcnNlL3JlZ2lzdHJhcnNlLnBhZ2Uuc2NzcyJ9 */"

/***/ }),

/***/ "./src/app/registrarse/registrarse.page.ts":
/*!*************************************************!*\
  !*** ./src/app/registrarse/registrarse.page.ts ***!
  \*************************************************/
/*! exports provided: RegistrarsePage */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "RegistrarsePage", function() { return RegistrarsePage; });
/* harmony import */ var tslib__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! tslib */ "./node_modules/tslib/tslib.es6.js");
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/core */ "./node_modules/@angular/core/fesm2015/core.js");
/* harmony import */ var _angular_router__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! @angular/router */ "./node_modules/@angular/router/fesm2015/router.js");
/* harmony import */ var _servicios_security_login_service__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ../servicios/security/login.service */ "./src/app/servicios/security/login.service.ts");
/* harmony import */ var _nebular_theme__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! @nebular/theme */ "./node_modules/@nebular/theme/fesm2015/index.js");





// import { MensajeComponent } from '../mensaje/mensaje.component';
let RegistrarsePage = class RegistrarsePage {
    constructor(router, dialogService, loginService, toastrService) {
        this.router = router;
        this.dialogService = dialogService;
        this.loginService = loginService;
        this.toastrService = toastrService;
        this.iUsuario = {
            cedula: '',
            tipo: '',
            componente: '',
            clave: '',
            correo: '',
        };
        this.lblcedula = 'Cédula del titular';
        this.loading = false;
        this.isHidden = true;
    }
    ngOnInit() {
    }
    open(dialog) {
        this.dialogService.open(dialog, { context: {
                title: 'Registro Ipsfa en linea',
            }, });
    }
    add(position, status, dialog) {
        this.loading = true;
        this.iUsuario.cedula = this.usuario;
        this.iUsuario.tipo = this.titular;
        this.iUsuario.componente = this.componente;
        this.iUsuario.clave = this.clave;
        this.iUsuario.correo = this.correo;
        if (this.clave != this.rclave) {
            status = 'danger';
            this.toastrService.show('Error en la clave deben ser identicas', `Ipsfa en línea`, { position, status });
            this.clave = '';
            this.rclave = '';
            this.loading = false;
            this.isHidden = false;
            return false;
        }
        this.loginService.makeUser(this.iUsuario).subscribe((data) => {
            status = 'success';
            //console.log(data);
            this.loading = false;
            this.isHidden = false;
            if (data.msj != undefined) {
                this.toastrService.show(data.msj, `Ipsfa en línea`, { position, status });
            }
            else {
                this.toastrService.show('Ok, proceso exitoso', `Ipsfa en línea`, { position, status });
                this.open(dialog);
            }
            //this.router.navigate(['login']);
        }, (error) => {
            //console.log(error.error.msj);
            this.loading = false;
            this.isHidden = false;
            var e = error.error.msj != undefined ? error.error.msj : 'Error de conexión';
            this.toastrService.show(e, `Ipsfa en línea`, { position, status });
        });
    }
    continuar() {
        this.router.navigate(['login']);
    }
    seleccionarTipo(valor) {
        switch (valor) {
            case "TIT":
                this.lblcedula = 'Cédula del titular';
                break;
            case "SOB":
                this.lblcedula = 'Cédula del sobreviviente';
                break;
            case "SOB-TIT":
                this.lblcedula = 'Cédula del titular - sobreviviente ';
                break;
            default:
                break;
        }
    }
};
RegistrarsePage = tslib__WEBPACK_IMPORTED_MODULE_0__["__decorate"]([
    Object(_angular_core__WEBPACK_IMPORTED_MODULE_1__["Component"])({
        selector: 'app-registrarse',
        template: __webpack_require__(/*! raw-loader!./registrarse.page.html */ "./node_modules/raw-loader/index.js!./src/app/registrarse/registrarse.page.html"),
        styles: [__webpack_require__(/*! ./registrarse.page.scss */ "./src/app/registrarse/registrarse.page.scss")]
    }),
    tslib__WEBPACK_IMPORTED_MODULE_0__["__metadata"]("design:paramtypes", [_angular_router__WEBPACK_IMPORTED_MODULE_2__["Router"],
        _nebular_theme__WEBPACK_IMPORTED_MODULE_4__["NbDialogService"],
        _servicios_security_login_service__WEBPACK_IMPORTED_MODULE_3__["LoginService"],
        _nebular_theme__WEBPACK_IMPORTED_MODULE_4__["NbToastrService"]])
], RegistrarsePage);



/***/ })

}]);
//# sourceMappingURL=registrarse-registrarse-module-es2015.js.map