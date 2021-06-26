<template>
  <div id="app">

    <div class="modal fade" id="modalAddNews" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">Добавить новость</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Закрыть"></button>
          </div>
          <div class="modal-body">
            <div class="input-group mb-3">
              <div class="input-group-prepend">
                <span class="input-group-text">Заголовок</span>
              </div>
              <input type="text" class="form-control" id="title" placeholder="Заголовок" aria-label="Username"
                     aria-describedby="basic-addon1">
            </div>
            <div class="input-group">
              <span class="input-group-text">Текст</span>
              <textarea class="form-control" id="text" aria-label="With textarea"></textarea>
            </div>
            <div class="input-group my-3">
              <label class="input-group-text" for="select-author">Автор</label>
              <select class="form-select" id="select-author" multiple>
              </select>
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" id="closeNews" class="btn btn-secondary" data-bs-dismiss="modal">Закрыть</button>
            <button type="button" class="btn btn-primary" @click="saveNews">Сохранить</button>
          </div>
        </div>
      </div>
    </div>

    <div class="modal fade" id="modalAddAuthor" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">Добавить автора</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Закрыть"></button>
          </div>
          <div class="modal-body">
            <div class="input-group mb-3">
              <div class="input-group-prepend">
                <span class="input-group-text">Имя</span>
              </div>
              <input type="text" class="form-control" id="first-name" placeholder="Имя" aria-label="Username"
                     aria-describedby="basic-addon1">
            </div>
            <div class="input-group mb-3">
              <div class="input-group-prepend">
                <span class="input-group-text">Фамилия</span>
              </div>
              <input type="text" class="form-control" id="last-name" placeholder="Фамилия" aria-label="Username"
                     aria-describedby="basic-addon1">
            </div>
            <div class="input-group mb-3">
              <div class="input-group-prepend">
                <span class="input-group-text">Отчество</span>
              </div>
              <input type="text" class="form-control" id="father-name" placeholder="Отчество" aria-label="Username"
                     aria-describedby="basic-addon1">
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" id="closeAuthor" class="btn btn-secondary" data-bs-dismiss="modal">Закрыть</button>
            <button type="button" class="btn btn-primary" @click="saveAuthor">Сохранить</button>
          </div>
        </div>
      </div>
    </div>


    <div class="container">
      <NewsTable
          v-bind:news="news"
          v-on:remove-news="removeNews"/>

    </div>
    <button type="button" class="m-2 btn btn-outline-primary btn-lg" data-bs-toggle="modal"
            data-bs-target="#modalAddNews">Добавить новость
    </button>
    <br>
    <button type="button" class="m-2 btn btn-outline-info btn-lg" data-bs-toggle="modal"
            data-bs-target="#modalAddAuthor">Добавить автора
    </button>
  </div>
</template>

<script>
import NewsTable from './components/NewsTable'
import axios from 'axios'

export default {
  name: 'App',
  components: {
    NewsTable,
  },
  data() {
    return {
      news: [],
      authors: []
    }
  },
  mounted() {
    this.getNews();
    this.getAuthors();
  },
  methods: {
    getAuthors() {
      axios.get('http://localhost:8081/api/author')
          .then(response => {
            this.authors = [];
            this.authors = response.data;
            this.updateSelectAuthor();
          });
    },
    getNews() {
      axios.get('http://localhost:8081/api/news')
          .then(response => {
            this.news = [];
            response.data.map(d => {
              this.news.push({
                id: d.News.id,
                text: d.News.text,
                title: d.News.title,
                created_at: d.News.created_at,
                authors: d.Authors
              })
            })
      });
    },
    removeNews(id) {
      axios.delete(`http://localhost:8081/api/news/${id}`)
          .then(() => {
            this.news = this.news.filter(n => n.id !== id);
        });
    },
    saveAuthor() {
      const firstName = document.getElementById('first-name').value;
      const lastName = document.getElementById('last-name').value;
      const fatherName = document.getElementById('father-name').value;
      if (firstName === '' || lastName === '') {
        console.log('Поля пустые');
        return;
      }
      axios.post('http://localhost:8081/api/author', {
        firstName: firstName,
        lastName: lastName,
        fatherName: fatherName
      }).then(() => {
        this.getAuthors();
        document.getElementById('closeAuthor').click();
      });
    },
    saveNews() {
      const title = document.getElementById('title').value;
      const text = document.getElementById('text').value;
      const authorsIds = this.getSelectValues(document.getElementById('select-author'));
      if (title === '' || text === '' || authorsIds.length === 0) {
        console.log('Поля пустые');
        return;
      }
      axios.post('http://localhost:8081/api/news', {
        title: title,
        text: text,
      }).then(response => {
        this.getNews();
        axios.post('http://localhost:8081/api/news/setnewstoauthor', {
          news_id: response.data,
          authors_ids: authorsIds,
        }).then(() => {
          this.getNews();
          document.getElementById('closeNews').click();
        });
      });
    },
    getSelectValues(select) {
      let result = [];
      let options = select && select.options;
      let opt;

      for (let i = 0, iLen = options.length; i < iLen; i++) {
        opt = options[i];

        if (opt.selected) {
          result.push(opt.value || opt.text);
        }
      }
      return result;
    },
    updateSelectAuthor() {
      const select = document.getElementById('select-author');
      this.removeOptions(select);
      this.authors.map(a => {
        let opt = document.createElement('option');
        opt.value = a.id;
        opt.innerHTML = `${a.lastName} ${a.firstName} ${a.fatherName}`;
        select.appendChild(opt);
      })
    },
    removeOptions(selectElement) {
      var i, L = selectElement.options.length - 1;
      for(i = L; i >= 0; i--) {
        selectElement.remove(i);
      }
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>