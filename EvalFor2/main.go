package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"time"
)

// MEDICINE ESTRUCTURA Y LO DEMAS

type Medicine struct {
	name            string
	manufacturer    string
	manufactureDate time.Time
	shelfLife       int
}

func NewMedicine(name, manufacturer string, manufactureDate time.Time, shelfLife int) (*Medicine, error) {
	m := &Medicine{}

	if err := m.SetName(name); err != nil {
		return m, err
	}
	if err := m.SetManufacturer(manufacturer); err != nil {
		return m, err
	}
	if err := m.SetManufactureDate(manufactureDate); err != nil {
		return m, err
	}
	if err := m.SetShelfLife(shelfLife); err != nil {
		return m, err
	}

	return m, nil
}

func (m Medicine) GetMed() (string, string, time.Time, int) {
	return m.name,
		m.manufacturer,
		m.manufactureDate,
		m.shelfLife
}

func (m *Medicine) SetName(name string) error {
	if name == "" {
		return errors.New("ERROR! El dato ingresado para nombre no es valido, no puede estar vacio")
	}
	m.name = name
	return nil
}

func (m Medicine) GetName() string {
	return m.name
}

func (m *Medicine) SetManufacturer(manufacturer string) error {
	if manufacturer == "" {
		return errors.New("ERROR! El dato ingresado para fabricante no es valido, no puede estar vacio")
	}
	m.manufacturer = manufacturer
	return nil
}
func (m Medicine) GetManufacturer() string {
	return m.manufacturer
}

func (m *Medicine) SetManufactureDate(date time.Time) error {
	if date.IsZero() {
		return errors.New("ERROR! El dato ingresado para fecha de fabricacion no es valido")
	}
	m.manufactureDate = date
	return nil
}
func (m Medicine) GetManufactureDate() time.Time {
	return m.manufactureDate
}

func (m *Medicine) SetShelfLife(months int) error {
	if months <= 0 {
		return errors.New("ERROR! El dato ingresado para vida util no es valido, debe ser un numero positivo")
	}
	m.shelfLife = months
	return nil
}
func (m Medicine) GetShelfLife() int {
	return m.shelfLife
}

func (m *Medicine) CalculateExpirationDate() time.Time {
	return m.manufactureDate.AddDate(0, m.shelfLife, 0)
}

// TABLET ESTRUCTURA Y LO DEMAS

type Tablet struct {
	Medicine
	dosePerTablet  float64
	isPrescription bool
}

func NewTablet(name, manufacturer string, mfgDate time.Time, shelfLife int, dose float64, prescription bool) (Tablet, error) {
	med, err := NewMedicine(name, manufacturer, mfgDate, shelfLife)
	if err != nil {
		return Tablet{}, err
	}
	t := Tablet{Medicine: *med, isPrescription: prescription}
	if err := t.SetDosePerTablet(dose); err != nil {
		return Tablet{}, err
	}
	return t, nil
}

func (t *Tablet) SetDosePerTablet(dose float64) error {
	if dose <= 0 {
		return errors.New("ERROR! El dato ingresado para la dosis no es valido, debe ser mayor a 0")
	}
	t.dosePerTablet = dose
	return nil
}
func (t *Tablet) GetDosePerTablet() float64 {
	return t.dosePerTablet
}
func (t *Tablet) SetIsPrescription(req bool) {
	t.isPrescription = req
}
func (t *Tablet) GetIsPrescription() bool {
	return t.isPrescription
}

func (t *Tablet) DisplayDetailsT() {
	receta := "No"
	if t.isPrescription {
		receta = "Sí"
	}
	fmt.Printf("Tableta %s | Fabricante: %s | Dosis: %.2f mg | Quiere Receta: %s\n", t.GetName(), t.GetManufacturer(), t.dosePerTablet, receta)
}

// SYRUP ESTRUCTURA Y LO DEMAS

type Syrup struct {
	Medicine
	volume float64
	flavor string
}

func NewSyrup(name, manufacturer string, mfgDate time.Time, shelfLife int, vol float64, flavor string) (Syrup, error) {
	med, err := NewMedicine(name, manufacturer, mfgDate, shelfLife)
	if err != nil {
		return Syrup{}, err
	}
	s := Syrup{Medicine: *med, flavor: flavor}
	if err := s.SetVolume(vol); err != nil {
		return Syrup{}, err
	}
	return s, nil
}

func (s *Syrup) SetVolume(vol float64) error {
	if vol <= 0 {
		return errors.New("ERROR! El dato ingresado para volumen no es valido, debe ser mayor a 0")
	}
	s.volume = vol
	return nil
}
func (s *Syrup) GetVolume() float64 {
	return s.volume
}
func (s *Syrup) SetFlavor(flavor string) {
	s.flavor = flavor
}
func (s *Syrup) GetFlavor() string {
	return s.flavor
}

func (s *Syrup) DisplayDetailsS() {
	fmt.Printf("Jarabe %s | Fabricante: %s | Volumen: %.2f ml | Sabor: %s\n",
		s.GetName(), s.GetManufacturer(), s.volume, s.flavor)
}

func main() {

	sc := bufio.NewScanner(os.Stdin)
	fmt.Println("Sistema de Gestión de Inventario de Medicamentos")

	date1 := time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC)

	t1, _ := NewTablet("Paracetamol", "PharmaCorp", date1, 24, 500.0, false)
	s1, _ := NewSyrup("Ambroxol", "BroncoResp", date1, 18, 120.0, "Fresa")

	tabletsInventory := []Tablet{t1}
	syrupsInventory := []Syrup{s1}

	var op int
	foriu := true

	for foriu {
		fmt.Println("\nMenú Principal")
		fmt.Println("Selecciona una opción: \n Opción 1: Agregar medicamento\n Opción 2: Ver inventarios\n Opción 3: Cambiar dosis de una tableta\n Opción 4: Ver fechas de caducidad\n Escriba 0 para salir del menú")

		fmt.Print("Escribe una opción: ")
		fmt.Scan(&op)

		switch op {
		case 0:
			fmt.Println("Salió del menú")
			foriu = false

		case 1:
			fmt.Println("Agregar medicamento")
			var tipo int
			fmt.Print("Ingrese 1 para Tableta o 2 para Jarabe ")
			fmt.Scan(&tipo)

			var nombre string
			var fabricante string
			var vidaUtil int

			fmt.Print("Ingrese el nombre: ")
			sc.Scan()
			nombre = sc.Text()
			fmt.Print("Ingrese el fabricante: ")
			sc.Scan()
			fabricante = sc.Text()
			fmt.Print("Ingrese la vida util en meses: ")
			fmt.Scan(&vidaUtil)

			if tipo == 1 {
				var dosis float64
				var recetanum int
				var receta bool
				fmt.Print("Ingrese la dosis en mg: ")
				fmt.Scan(&dosis)
				fmt.Print("Necesita receta? 1 es si, 0 es no: ")
				fmt.Scan(&recetanum)
				if recetanum == 1 {
					receta = true
				} else {
					receta = false
				}

				nuevaTablet, err := NewTablet(nombre, fabricante, date1, vidaUtil, dosis, receta)
				if err == nil {
					tabletsInventory = append(tabletsInventory, nuevaTablet)
					fmt.Println("La tableta se agrego con exito")
				} else {
					fmt.Println("Error al agregar la Tableta:", err)
				}
			} else if tipo == 2 {
				var volumen float64
				var sabor string
				fmt.Print("Ingrese el volumen del Jarabe en ml: ")
				fmt.Scan(&volumen)
				fmt.Print("Ingrese el sabor del Jarabe: ")
				sc.Scan()
				sabor = sc.Text()

				nuevoSyrup, err := NewSyrup(nombre, fabricante, date1, vidaUtil, volumen, sabor)
				if err == nil {
					syrupsInventory = append(syrupsInventory, nuevoSyrup)
					fmt.Println("El Jarabe se agrego con exito")
				} else {
					fmt.Println("Error al agregar el Jarabe:", err)
				}
			} else {
				fmt.Println("ERROR, Ingrese una opcion valida")
			}

			foriu = false

		case 2:
			fmt.Println("Lista de Medicamentos")
			fmt.Println("Tabletas")
			if len(tabletsInventory) == 0 {
				fmt.Println("El inventario de tabletas esta vacio")
			} else {
				for i, t := range tabletsInventory {
					fmt.Printf("ID: %d ", i)
					t.DisplayDetailsT()
				}
			}

			fmt.Println("Jarabes")
			if len(syrupsInventory) == 0 {
				fmt.Println("El inventario de Jarabes esta vacio")
			} else {
				for i, s := range syrupsInventory {
					fmt.Printf("ID: %d ", i)
					s.DisplayDetailsS()
				}
			}

			foriu = false

		case 3:
			fmt.Println("Cambiar dosis de Tableta")
			var id int
			fmt.Print("Ingrese el ID de la tableta que desea actualizar: ")
			fmt.Scan(&id)

			if id >= 0 && id < len(tabletsInventory) {
				var nuevaDosis float64
				fmt.Print("Ingrese la nueva dosis en mg: ")
				fmt.Scan(&nuevaDosis)
				err := tabletsInventory[id].SetDosePerTablet(nuevaDosis)
				if err == nil {
					fmt.Println("La tableta se actualizo con exito")
					fmt.Printf("La nueva dosis es: %.2f mg\n", tabletsInventory[id].GetDosePerTablet())
				} else {
					fmt.Println("Error:", err)
				}
			} else {
				fmt.Println("ERROR, el ID no es valido")
			}

			foriu = false

		case 4:
			fmt.Println("Fechas de Caducidad de los Medicamentos")
			for _, t := range tabletsInventory {
				exp := t.CalculateExpirationDate()
				fmt.Printf("Tableta: %s | Fab: %s | Caducidad: %s\n", t.GetName(), t.GetManufactureDate().Format("2006-01-02"), exp.Format("2006-01-02"))
			}
			for _, s := range syrupsInventory {
				exp := s.CalculateExpirationDate()
				fmt.Printf("Jarabe: %s | Fab: %s | Caducidad: %s\n", s.GetName(), s.GetManufactureDate().Format("2006-01-02"), exp.Format("2006-01-02"))
			}

			foriu = false
		}
	}
}
