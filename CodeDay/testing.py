
import requests
from re import DEBUG, sub
from flask import Flask, render_template, request, redirect, send_file, url_for, jsonify
from werkzeug.utils import secure_filename, send_from_directory
import os
import subprocess
from flask_cors import CORS

API_key = "64b38d3f58717f4e8c8c6a8033addac1"
 
# This stores the url
base_url = "http://api.openweathermap.org/data/2.5/weather?"
base_name_url = "http://history.openweathermap.org/data/2.5/history/city?q="
# Example "http://api.openweathermap.org/geo/1.0/direct?q=London&limit=5&appid={API key}"
base_url_geocoding = "http://api.openweathermap.org/geo/1.0/direct?q="

city_name = input("Enter a city name : ")
 
# final_url = base_url + "appid=" + API_key + "&id=" + city_id
# final_url_geocoding = base_url_geocoding + city_name + "&appid=" + API_key


final_url_geocoding = base_url + "appid=" + API_key + "&q=" + city_name

weather_data = requests.get(final_url_geocoding).json()



print(weather_data)


