package adventofcode.year2019;

import java.util.ArrayList;

import org.apache.commons.lang3.StringUtils;
import org.junit.Test;

import adventofcode.BaseTest;

public class Day8 extends BaseTest {

	@Test
	public void singleCheck() {

	}

	@Test
	public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2019/day8/input.txt");
		int count = 0;
		String layer = "";
		for (String input : data) {
			count = input.length();
			layer = findLayer(input, 25, 6);
		}
		System.out.println(getSilverResult(layer));
	}

	@Test
	public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2019/day8/input.txt");
		int count = 0;
		ArrayList<String> layersList = null;
		for (String input : data) {
			count = input.length();
			layersList = readLayers(input, 25, 6);
		}

		String finalLayer = getFinalLayer(layersList);
		System.out.println(finalLayer);
//		System.out.println(getSilverResult(layer));
		for (int i=0; i< finalLayer.length(); i= i+ 25) {
//			System.out.println(input.substring(i, i + wide*tall));
			System.out.println(finalLayer.substring(i, i + 25));
		}
	}


	public String getFinalLayer(ArrayList<String> layers){
		int length = layers.get(0).length();
		String finalLayer = "";
		System.out.println("layers size: " + layers.size());
		for (int i = 0; i < length; i++) {
			boolean found = false;
			for (int j=0; j< layers.size(); j++){
				char currentChar = layers.get(j).charAt(i);
				if (currentChar != '2' && !found){
					found = true;
//					System.out.println("i-> " + i + " j-> "+ j );
					finalLayer = finalLayer + currentChar;
				}
			}
			if (!found){
				finalLayer = finalLayer + '2';
			}
		}
		return finalLayer;
	}
	public ArrayList<String> readLayers(final String input, int wide, int tall){
		ArrayList<String> layers = new ArrayList<String>();
		for (int i=0; i< input.length(); i= i+ wide*tall) {
//			System.out.println(input.substring(i, i + wide*tall));
			layers.add(input.substring(i, i + wide*tall));
		}
		return layers;
	}


	public String findLayer(final String input, int wide, int tall){
		String minLayer = "";
		int minMatch = Integer.MAX_VALUE;
		for (int i=0; i< input.length(); i= i+ wide*tall) {
			String currentLayer = input.substring(i, i + wide*tall);
			int currentMatch = StringUtils.countMatches(currentLayer, "0");
			if (currentMatch < minMatch){
				minMatch = currentMatch;
				minLayer = currentLayer;
//				System.out.println(i +" -> " + currentMatch+ " -> "+ currentLayer);
			}
		}
		return minLayer;
	}

	public int getSilverResult(final String layer){
		return StringUtils.countMatches(layer, "1") * StringUtils.countMatches(layer , "2");
	}
}
