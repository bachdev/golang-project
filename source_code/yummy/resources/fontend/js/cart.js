$(document).ready(function () {

    // san pham
    $('[data-toggle="filter-display"]').each(function () {
      $(this).click(function () {
        displayTarget = $(this).data('target');
        $('.product-group, .sub-group').removeClass('active');
        $(this).addClass('active');
        $('.col-item').hide();
        $('.col-item' + displayTarget).fadeIn(600);
      });
    });
    $('[data-toggle="display-all"]').on('click', function () {
      $('.product-group, .sub-group').removeClass('active');
      $('.col-item').fadeIn(600);
    });
    //chi tiet
    $('.qty-decrease').each(function () {
      $(this).on('click', function () {
        var updateQty = Number($(this).next('input').attr('value'));
        if (updateQty > 1) {
          updateQty--;
        }
        $(this).next('input').attr('value', updateQty);
        $(this).next('input').val(updateQty);
        updateSubtotal(this, updateQty);
        updateTotal();
      });
    });
    $('.qty-increase').each(function () {
      $(this).on('click', function () {
        var updateQty = Number($(this).prev('input').attr('value'));
        updateQty++;
        $(this).prev('input').attr('value', updateQty);
        $(this).prev('input').val(updateQty);
        updateSubtotal(this, updateQty);
        updateTotal();
      });
    });

    //gio hang
    $('td.product-price').each(function () {
      var item = $(this).text();
      sum += Number($(this).text());
      discount = sum * discountRate / 100;
      total = sum - discount;
      var qty = $(this).next('td').find('input').attr('value');
      var subTotal = (Number(item) * Number(qty)).toLocaleString('vi');
      var num = Number(item).toLocaleString('vi');
      $(this).text(num);
      $(this).siblings('.product-subtotal').text(subTotal);
      $('.cart-subtotal td').text(sum.toLocaleString('vi'));
      $('.amount').text(total.toLocaleString('vi'));
    });

    //thanhtoan
    $('.product-quantity input').each(function () {
      $(this).change(function () {
        var currentQty = $(this).val();
        $(this).attr('value', currentQty);
        updateSubtotal(this, currentQty);
        updateTotal();
      });
    });
    $('.product-remove button').each(function () {
      $(this).on('click', function () {
        $(this).parents('.cart-item').remove();
        updateTotal();
      });
    });

    $('.apply-coupon').click(function () {
      discountRate = $(this).prev('input').val();
      $('.discount-rate').text(discountRate + '%');
      $('.cart-discount td').text(discount.toLocaleString('vi'));
      updateTotal();
    });

});





var sum = 0, total, discount, discountRate=0;
// var $navMenuCont;
function updateSubtotal(currentItem, currentQty) {
var item = $(currentItem).parents('.product-quantity').siblings('.product-price').text().replace(/\./g, "");
var subTotal = (Number(item) * currentQty).toLocaleString('vi');
$(currentItem).parents('.product-quantity').siblings('.product-subtotal').text(subTotal);
}
function updateTotal() {
sum = 0;
$('td.product-subtotal').each(function () {
  sum += Number($(this).text().replace(/\./g, ""));
});
discount = sum * discountRate / 100;
total = sum - discount;
$('.cart-subtotal td').text(sum.toLocaleString('vi'));
$('.cart-discount td').text(discount.toLocaleString('vi'));
$('.amount').text(total.toLocaleString('vi'));
}
