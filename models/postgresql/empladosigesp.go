package models

import (
	"encoding/json"
	"strconv"

	"github.com/elpoloxrodriguez/esb-inea/sys"
)

type BuscarEmpleadoSigesp struct {
	ID        int    `json:"id,omitempty"`
	Codper    string `json:"codper,omitempty"`
	Nomper    string `json:"nomper,omitempty"`
	Apeper    string `json:"apeper,omitempty"`
	Fecnacper string `json:"fecnacper,omitempty"`
}

//Listar el Empleado SIGESP de PostgreSql
func (u *BuscarEmpleadoSigesp) MostrarEmpleadoSIGESP() (jSon []byte, err error) {
	rows, err := sys.PostgreSQL.Query(
		`select per.codper,
    per.nomper,
    per.apeper,
    per.fecnacper,
    per.edocivper,
    per.sexper,
    per.rifper,
    per.fecingper,
    per.anoservpreper as anosreconocidos,
    (case     when asi.denasicar is null and nom.codnom = '0006' then 'JUBILADO'
        when asi.denasicar is null and nom.codnom = '0007' then 'PENSIONADOS'
        when asi.denasicar = 'Sin Asignación de Cargo' then car.descar
        else asi.denasicar end),
    ger.denger,
    (case     when ban.nomban is null and substr(nom.codcueban, 1,4) = '0102' then 'VENEZUELA (Solo Uso Nomina)'
        when ban.nomban is null and substr(nom.codcueban, 1,4) = '0177' then 'BANFANB (NOMINA CIVIL)'
        else ban.nomban end) as nombrebanco,
    nom.codcueban,
    (case     when nom.codnom = '0001' then 'Empleados Fijos'
        when nom.codnom = '0002' then 'Empleados Contratados'
        when nom.codnom = '0003' then 'Obreros Fijos'
        when nom.codnom = '0004' then 'Obreros Contratos'
        when nom.codnom = '0006' then 'Jubilados'
        else 'Pensionados' end) tiponomina
        from    sno_personal as per
        inner join sno_personalnomina as nom
            on per.codper = nom.codper
            and nom.codnom in ('0001','0002','0003','0006','0007')
            and nom.staper in ('1','2')
    left join srh_gerencia as ger
        on nom.codemp = ger.codemp
        and per.codger = ger.codger
    left join scb_banco as ban
        on ban.codemp = ger.codemp
        and ban.codban = nom.codban
    left join sno_asignacioncargo as asi
        on ban.codemp = asi.codemp
        and nom.codnom = asi.codnom
        and nom.codasicar = asi.codasicar
    left join sno_cargo as car
        on car.codemp = nom.codemp
        and car.codnom = nom.codnom
        and car.codcar = nom.codcar
		where     per.codper = '015132444'
		order by nom.codnom, 1`,
	)
	if err != nil {
		panic(err.Error())
	}
	var lst []interface{}
	for rows.Next() {
		var mysqldata BuscarEmpleadoSigesp
		var id, codper, nomper, apeper, fecnacper string
		err = rows.Scan(&id, &codper, &nomper, &apeper, &fecnacper)
		if err != nil {
			panic(err.Error())
		}
		mysqldata.ID, _ = strconv.Atoi(id)
		mysqldata.Codper = codper
		mysqldata.Nomper = nomper
		mysqldata.Apeper = apeper
		mysqldata.Fecnacper = fecnacper
		lst = append(lst, mysqldata)
	}
	jSon, err = json.Marshal(lst)
	return
}
