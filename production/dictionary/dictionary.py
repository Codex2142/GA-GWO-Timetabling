import os
import pandas as pd
import numpy as np
from collections import defaultdict

# Path relatif ke file dictionary.py
MAIN_DIR = os.path.join(os.path.dirname(__file__), "../data/csv")

# ===============================
# Load semua CSV
# ===============================
guru_df = pd.read_csv(os.path.join(MAIN_DIR, "guru.csv"))
mapel_df = pd.read_csv(os.path.join(MAIN_DIR, "mapel.csv"))
kelas_df = pd.read_csv(os.path.join(MAIN_DIR, "kelas.csv"))
slot_df = pd.read_csv(os.path.join(MAIN_DIR, "slot.csv"))
relasi_guru_mapel_df = pd.read_csv(os.path.join(MAIN_DIR, "relasi_guru_mapel.csv"))
wali_kelas_df = pd.read_csv(os.path.join(MAIN_DIR, "wali_kelas.csv"))

# ===============================
# Mapping hari -> id
# ===============================
def hariId():
    return {
        "Senin": 1,
        "Selasa": 2,
        "Rabu": 3,
        "Kamis": 4,
        "Jumat": 5,
    }

# ===============================
# Slot per hari
# ===============================
def slotPerHari():
    hd = hariId()
    slotPerHariDict = slot_df['hari'].value_counts().to_dict()
    slotPerHariDict = {hd[k]: v for k, v in slotPerHariDict.items()}
    slotPerHari = np.array([slotPerHariDict[i] for i in range(1,6)])
    return slotPerHari

# ===============================
# Kelas -> tingkatan
# ===============================
def kelasTingkatan():
    return dict(zip(kelas_df['kelas_id'], kelas_df['tingkatan']))

def kelasIds():
    return np.array(sorted(kelasTingkatan().keys()))

def jumlahKelas():
    return len(kelasIds())

# ===============================
# Mapel info
# ===============================
def jamPerMingguMapel():
    return dict(zip(mapel_df['mapel_id'], mapel_df['jam_per_minggu']))

def mapelIds():
    return np.array(list(jamPerMingguMapel().keys()))

def jamMapel():
    return np.array(list(jamPerMingguMapel().values()))

# ===============================
# MGMP Mapel
# ===============================
def mgmpMapel():
    hd = hariId()
    mgmp = dict(zip(mapel_df['mapel_id'], mapel_df['MGMP']))
    return {mapel_id: hd[hari] for mapel_id, hari in mgmp.items()}

# ===============================
# Batas slot siang
# ===============================
def batasSiang():
    return np.array([5,5,4,5,4])

# ===============================
# Batas MGMP per hari
# ===============================
def batasMGMP():
    return np.array([2,2,2,2,1])

# ===============================
# Guru valid per mapel + tingkatan
# ===============================
def guruValidMapel():
    d = defaultdict(list)
    for row in relasi_guru_mapel_df.itertuples():
        key = (row.mapel_id, row.tingkatan)
        d[key].append(row.guru_id)
    return dict(d)

# ===============================
# Total durasi guru
# ===============================
def maxJamGuru():
    return relasi_guru_mapel_df.groupby("guru_id")["durasi"].sum().to_dict()

# ===============================
# Wali kelas mapping
# ===============================
def waliKelas():
    return dict(zip(wali_kelas_df['guru_id'], wali_kelas_df['kelas_id']))

# ===============================
# Slot -> hari
# ===============================
def slotHari():
    hd = hariId()
    return np.array([hd[row.hari] for row in slot_df.itertuples()])

# ===============================
# Total slot per kelas dan total gen
# ===============================
def slotPerKelas():
    return len(slotHari())

def totalSlot():
    return jumlahKelas() * slotPerKelas()

# ===============================
# Fitness weight
# ===============================
def fitnessWeight():
    return {
        "guruBentrok": 100,
        "distribusiMapel": 20,
        "konsistensiGuruMapel": 20,
        "durasiGuru": 10,
        "waktuMGMP": 5,
        "mapelSiang": 5,
        "cekWaliKelas": 5
    }

# ===============================
# Blok jam mapel
# ===============================
def blokDistribusi(jam):
    if jam == 2: return [2]
    if jam == 3: return [3]
    if jam == 4: return [2,2]
    if jam == 5: return [2,3]
    return [jam]

def blokMapel():
    blok = {}
    for mapel_id, jam in jamPerMingguMapel().items():
        blok[mapel_id] = blokDistribusi(jam)
    return blok

# ===============================
# Slot awal dan akhir hari
# ===============================
def slotAwalHari():
    return np.concatenate(([0], np.cumsum(slotPerHari())[:-1]))

def slotAkhirHari():
    return np.cumsum(slotPerHari())